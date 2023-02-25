package workspace

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/epermission"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionEstimate(p *Params) (*FullEstimate, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fe, err := s.LoadEstimate(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}
	switch p.Act {
	case action.ActUpdate:
		return estimateUpdate(p, fe)
	case action.ActChildAdd:
		return estimateStoryAdd(p, fe)
	case action.ActChildUpdate:
		return estimateStoryUpdate(p, fe)
	case action.ActChildStatus:
		return estimateStoryStatus(p, fe)
	case action.ActVote:
		return estimateStoryVote(p, fe)
	case action.ActChildRemove:
		return estimateStoryRemove(p, fe)
	case action.ActMemberUpdate:
		return estimateMemberUpdate(p, fe)
	case action.ActMemberRemove:
		return estimateMemberRemove(p, fe)
	case action.ActMemberSelf:
		return estimateUpdateSelf(p, fe)
	case action.ActComment:
		return estimateComment(p, fe)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func estimateUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
		return nil, "", "", errors.New("you do not have permission to edit this estimate")
	}
	tgt := fe.Estimate.Clone()
	tgt.Title = p.Frm.GetStringOpt("title")
	if tgt.Title == "" {
		return nil, "", "", errors.New("title may not be empty")
	}
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tx, err := p.Svc.db.StartTransaction(p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	tgt.Slug = p.Svc.e.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.eh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Choices = util.StringSplitAndTrim(p.Frm.GetStringOpt("choices"), ",")
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fe.Estimate.Diff(tgt)) > 0
	permsChanged := len(fe.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fe, MsgNoChangesNeeded, fe.Estimate.PublicWebPath(), nil
	}
	if modelChanged {
		model, err := p.Svc.SaveEstimate(p.Ctx, tgt, fe.Self.UserID, tx, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = updateTeam(
			util.KeyEstimate, fe.Estimate.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fe.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyEstimate, fe.Estimate.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fe.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fe.Estimate = model
		err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActUpdate, model, &fe.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(util.KeyEstimate, model.TeamID, model.SprintID, model, &fe.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if err := p.Svc.ep.DeleteWhere(p.Ctx, tx, "estimate_id = $1", len(fe.Permissions), p.Logger, tgt.ID); err != nil {
			return nil, "", "", err
		}
		newPerms := make(epermission.EstimatePermissions, 0, len(perms))
		for _, x := range perms {
			newPerms = append(newPerms, &epermission.EstimatePermission{EstimateID: tgt.ID, Key: x.Key, Value: x.Value})
		}
		if err = p.Svc.ep.Save(p.Ctx, tx, p.Logger, newPerms...); err != nil {
			return nil, "", "", err
		}
		err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActPermissions, perms, &fe.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Estimate updated", fe.Estimate.PublicWebPath(), nil
}

func estimateStoryAdd(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	title := strings.TrimSpace(p.Frm.GetStringOpt("title"))
	if title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	st := &story.Story{
		ID: util.UUID(), EstimateID: fe.Estimate.ID, Idx: len(fe.Stories), UserID: fe.Self.UserID, Title: title, Status: enum.SessionStatusNew, Created: time.Now(),
	}
	err := p.Svc.st.Create(p.Ctx, nil, p.Logger, st)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildAdd, st, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story added", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	if curr.UserID != fe.Self.UserID && (!fe.Admin()) {
		return nil, "", "", errors.New("you do not have permission to edit this story")
	}
	st := curr.Clone()
	st.Title = strings.TrimSpace(p.Frm.GetStringOpt("title"))
	if st.Title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	if len(curr.Diff(st)) == 0 {
		return fe, MsgNoChangesNeeded, st.PublicWebPath(fe.Estimate.Slug), nil
	}
	err := p.Svc.st.Update(p.Ctx, nil, st, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildUpdate, st, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story saved", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryStatus(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	statusStr := strings.TrimSpace(p.Frm.GetStringOpt("status"))
	if statusStr == "" {
		return nil, "", "", errors.New("must provide [status]")
	}
	status := enum.SessionStatus(statusStr)
	st := &story.Story{ID: *id, EstimateID: fe.Estimate.ID, Idx: curr.Idx, UserID: fe.Self.UserID, Title: curr.Title, Status: status, Created: curr.Created}
	err := p.Svc.st.Update(p.Ctx, nil, st, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save new status for story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildStatus, st, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story status updated", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryVote(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	currStory := fe.Stories.Get(*id)
	if currStory == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	voteStr := strings.TrimSpace(p.Frm.GetStringOpt("vote"))
	if voteStr == "" {
		return nil, "", "", errors.New("must provide [vote]")
	}
	if !slices.Contains(fe.Estimate.Choices, voteStr) {
		return nil, "", "", errors.Errorf("vote choice [%s] is not one of the valid choices [%s]", voteStr, strings.Join(fe.Estimate.Choices, ", "))
	}
	v := fe.Votes.Get(*id, fe.Self.UserID)
	if v == nil {
		v = &vote.Vote{StoryID: *id, UserID: fe.Self.UserID, Created: time.Now()}
	}
	v.Choice = voteStr
	err := p.Svc.v.Save(p.Ctx, nil, p.Logger, v)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save vote for story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActVote, v, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Vote recorded", currStory.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryRemove(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	id, _ := p.Frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	if curr.UserID != fe.Self.UserID && (!fe.Admin()) {
		return nil, "", "", errors.New("you do not have permission to remove this story")
	}
	for _, v := range fe.Votes.GetByStoryIDs(*id) {
		err := p.Svc.v.Delete(p.Ctx, nil, v.StoryID, v.UserID, p.Logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "unable to delete vote")
		}
	}
	err := p.Svc.st.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete story")
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActChildRemove, id, &fe.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Story deleted", fe.Estimate.PublicWebPath(), nil
}

func estimateMemberUpdate(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.em.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberUpdate, curr, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgMemberUpdated, fe.Estimate.PublicWebPath(), nil
}

func estimateMemberRemove(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	if !fe.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fe.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	err := p.Svc.em.Delete(p.Ctx, nil, curr.EstimateID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberRemove, userID, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgMemberRemoved, fe.Estimate.PublicWebPath(), nil
}

func estimateUpdateSelf(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fe.Self.Name && picture == fe.Self.Picture {
		return fe, MsgNoChangesNeeded, fe.Estimate.PublicWebPath(), nil
	}

	fe.Self.Picture = picture
	fe.Self.Name = name
	err := p.Svc.em.Update(p.Ctx, nil, fe.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fe.Self.UserID, "name": name, "role": fe.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActMemberUpdate, arg, &fe.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgProfileEdited, fe.Estimate.PublicWebPath(), nil
}

func estimateComment(p *Params, fe *FullEstimate) (*FullEstimate, string, string, error) {
	c, u, err := commentFromForm(p.Frm, fe.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceEstimate:
		if c.ModelID != fe.Estimate.ID {
			return nil, "", "", errors.New("this comment refers to a different estimate")
		}
	case enum.ModelServiceStory:
		if curr := fe.Stories.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a story that isn't part of this estimate")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceEstimate, fe.Estimate.ID, c, &fe.Self.UserID, fe.Estimate.TeamID, fe.Estimate.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fe, MsgCommentAdded, fe.Estimate.PublicWebPath() + u, nil
}
