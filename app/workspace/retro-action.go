package workspace

import (
	"time"

	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/retro/rpermission"
	"github.com/kyleu/rituals/app/sprint"
	"github.com/kyleu/rituals/app/team"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionRetro(p *Params) (*FullRetro, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	fr, err := p.Svc.LoadRetro(lp, func() (team.Teams, error) {
		return p.Svc.t.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	}, func() (sprint.Sprints, error) {
		return p.Svc.s.GetByMember(p.Ctx, nil, p.Profile.ID, nil, p.Logger)
	})
	if err != nil {
		return nil, "", "", err
	}

	switch p.Act {
	case action.ActUpdate:
		return retroUpdate(p, fr)
	case action.ActChildAdd:
		return retroFeedbackAdd(p, fr)
	case action.ActChildUpdate:
		return retroFeedbackUpdate(p, fr)
	case action.ActChildRemove:
		return retroFeedbackRemove(p, fr)
	case action.ActMemberUpdate:
		return retroMemberUpdate(p, fr)
	case action.ActMemberRemove:
		return retroMemberRemove(p, fr)
	case action.ActMemberSelf:
		return retroUpdateSelf(p, fr)
	case action.ActComment:
		return retroComment(p, fr)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func retroUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
		return nil, "", "", errors.New("you do not have permission to update this retro")
	}
	tgt := fr.Retro.Clone()
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
	tgt.Slug = p.Svc.r.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.rh, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Categories = util.StringSplitAndTrim(p.Frm.GetStringOpt("categories"), ",")
	tgt.TeamID, _ = p.Frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = p.Frm.GetUUID(util.KeySprint, true)
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(fr.Retro.Diff(tgt)) > 0
	permsChanged := len(fr.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	if len(fr.Retro.Diff(tgt)) == 0 && len(fr.Permissions.ToPermissions().Diff(perms)) == 0 {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	if modelChanged {
		model, err := p.Svc.SaveRetro(p.Ctx, tgt, fr.Self.UserID, tx, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = updateTeam(
			util.KeyRetro, fr.Retro.TeamID, model.TeamID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fr.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		err = updateSprint(
			util.KeyRetro, fr.Retro.SprintID, model.SprintID, model.ID, model.TitleString(), model.PublicWebPath(), model.IconSafe(), fr.Self.UserID, p,
		)
		if err != nil {
			return nil, "", "", err
		}
		fr.Retro = model
		err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActUpdate, model, &fr.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		err = sendTeamSprintUpdates(util.KeyRetro, model.TeamID, model.SprintID, model, &fr.Self.UserID, p.Svc, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if err := p.Svc.rp.DeleteWhere(p.Ctx, tx, "retro_id = $1", len(fr.Permissions), p.Logger, tgt.ID); err != nil {
			return nil, "", "", err
		}
		newPerms := make(rpermission.RetroPermissions, 0, len(perms))
		for _, x := range perms {
			newPerms = append(newPerms, &rpermission.RetroPermission{RetroID: tgt.ID, Key: x.Key, Value: x.Value})
		}
		if err = p.Svc.rp.Save(p.Ctx, tx, p.Logger, newPerms...); err != nil {
			return nil, "", "", err
		}
		err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActPermissions, perms, &fr.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Retro updated", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackAdd(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	category := p.Frm.GetStringOpt("category")
	content := p.Frm.GetStringOpt("content")
	if content == "" {
		return nil, "", "", errors.New("must provide [content]")
	}
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: util.UUID(), RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: time.Now(),
	}
	err := p.Svc.f.Create(p.Ctx, nil, p.Logger, f)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildAdd, f, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback added", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [feedbackID]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	if curr.UserID != fr.Self.UserID && (!fr.Admin()) {
		return nil, "", "", errors.New("you do not have permission to update this feedback")
	}
	f := curr.Clone()
	f.Category = p.Frm.GetStringOpt("category")
	f.Content = p.Frm.GetStringOpt("content")
	f.HTML = util.ToHTML(f.Content, true)
	if len(curr.Diff(f)) == 0 {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}
	err := p.Svc.f.Update(p.Ctx, nil, f, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildUpdate, f, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback saved", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	id, _ := p.Frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [feedbackID]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	if curr.UserID != fr.Self.UserID && (!fr.Admin()) {
		return nil, "", "", errors.New("you do not have permission to remove this feedback")
	}
	err := p.Svc.f.Delete(p.Ctx, nil, *id, p.Logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete feedback")
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActChildRemove, id, &fr.Self.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Feedback deleted", fr.Retro.PublicWebPath(), nil
}

func retroMemberUpdate(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
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
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.rm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, curr, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgMemberUpdated, fr.Retro.PublicWebPath(), nil
}

func retroMemberRemove(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	if !fr.Admin() {
		return nil, "", "", errors.New("you do not have permission to remove this member")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	if *userID == fr.Self.UserID {
		return nil, "", "", errors.New("you can't remove yourself")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	err := p.Svc.rm.Delete(p.Ctx, nil, curr.RetroID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberRemove, userID, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgMemberRemoved, fr.Retro.PublicWebPath(), nil
}

func retroUpdateSelf(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	name := p.Frm.GetStringOpt("name")
	choice := p.Frm.GetStringOpt("choice")
	picture := p.Frm.GetStringOpt("picture")

	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	if name == fr.Self.Name && picture == fr.Self.Picture && choice != KeyGlobal {
		return fr, MsgNoChangesNeeded, fr.Retro.PublicWebPath(), nil
	}

	fr.Self.Picture = picture
	fr.Self.Name = name
	err := p.Svc.rm.Update(p.Ctx, nil, fr.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		err = p.Svc.SetName(p.Ctx, p.Profile.ID, name, picture, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	arg := util.ValueMap{"userID": fr.Self.UserID, "name": name, "role": fr.Self.Role}
	if picture != "" {
		arg["picture"] = picture
	}
	err = p.Svc.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, arg, &fr.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgProfileEdited, fr.Retro.PublicWebPath(), nil
}

func retroComment(p *Params, fr *FullRetro) (*FullRetro, string, string, error) {
	c, u, err := commentFromForm(p.Frm, fr.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceRetro:
		if c.ModelID != fr.Retro.ID {
			return nil, "", "", errors.New("this comment refers to a different retro")
		}
	case enum.ModelServiceFeedback:
		if curr := fr.Feedbacks.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to a feedback that isn't part of this retro")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = sendComment(enum.ModelServiceRetro, fr.Retro.ID, c, &fr.Self.UserID, fr.Retro.TeamID, fr.Retro.SprintID, p.Svc.send, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return fr, MsgCommentAdded, fr.Retro.PublicWebPath() + u, nil
}
