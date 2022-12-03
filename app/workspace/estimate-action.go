package workspace

import (
	"context"
	"github.com/kyleu/rituals/app/action"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"

	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/estimate/story/vote"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionEstimate(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullEstimate, string, string, error) {
	fe, err := s.LoadEstimate(ctx, slug, userID, "", nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	switch act {
	case action.ActUpdate:
		return estimateUpdate(ctx, fe, userID, frm, slug, s, logger)
	case action.ActStoryAdd:
		return estimateStoryAdd(ctx, fe, userID, frm, s, logger)
	case action.ActStoryUpdate:
		return estimateStoryUpdate(ctx, fe, userID, frm, s, logger)
	case action.ActStoryStatus:
		return estimateStoryStatus(ctx, fe, userID, frm, s, logger)
	case action.ActVote:
		return estimateStoryVote(ctx, fe, userID, frm, s, logger)
	case action.ActStoryRemove:
		return estimateStoryRemove(ctx, fe, userID, frm, s, logger)
	case action.ActMemberUpdate:
		return estimateMemberUpdate(ctx, fe, frm, s, logger)
	case action.ActMemberRemove:
		return estimateMemberRemove(ctx, fe, frm, s, logger)
	case action.ActMemberSelf:
		return estimateUpdateSelf(ctx, fe, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func estimateUpdate(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, slug string, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	tgt := fe.Estimate.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.r.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.rh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Choices = util.StringSplitAndTrim(frm.GetStringOpt("choices"), ",")
	tgt.TeamID, _ = frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = frm.GetUUID(util.KeySprint, true)
	model, err := s.SaveEstimate(ctx, tgt, userID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	fe.Estimate = model
	return fe, "Estimate saved", model.PublicWebPath(), nil
}

func estimateStoryAdd(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	title := strings.TrimSpace(frm.GetStringOpt("title"))
	if title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	st := &story.Story{
		ID: util.UUID(), EstimateID: fe.Estimate.ID, Idx: len(fe.Stories), UserID: userID, Title: title, Status: enum.SessionStatusNew, Created: time.Now(),
	}
	err := s.st.Create(ctx, nil, logger, st)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	return fe, "Story added", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryUpdate(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	id, _ := frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	title := strings.TrimSpace(frm.GetStringOpt("title"))
	if title == "" {
		return nil, "", "", errors.New("must provide [title]")
	}
	st := &story.Story{ID: *id, EstimateID: fe.Estimate.ID, Idx: curr.Idx, UserID: userID, Title: title, Status: curr.Status, Created: curr.Created}
	err := s.st.Update(ctx, nil, st, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited story")
	}
	return fe, "Story saved", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryStatus(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	id, _ := frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	statusStr := strings.TrimSpace(frm.GetStringOpt("status"))
	if statusStr == "" {
		return nil, "", "", errors.New("must provide [status]")
	}
	status := enum.SessionStatus(statusStr)
	st := &story.Story{ID: *id, EstimateID: fe.Estimate.ID, Idx: curr.Idx, UserID: userID, Title: curr.Title, Status: status, Created: curr.Created}
	err := s.st.Update(ctx, nil, st, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save new status for story")
	}
	return fe, "Story status updated", st.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryVote(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	id, _ := frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [storyID]")
	}
	currStory := fe.Stories.Get(*id)
	if currStory == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	voteStr := strings.TrimSpace(frm.GetStringOpt("vote"))
	if voteStr == "" {
		return nil, "", "", errors.New("must provide [vote]")
	}
	if !slices.Contains(fe.Estimate.Choices, voteStr) {
		return nil, "", "", errors.Errorf("vote choice [%s] is not one of the valid choices [%s]", voteStr, strings.Join(fe.Estimate.Choices, ", "))
	}
	v := fe.Votes.Get(*id, userID)
	if v == nil {
		v = &vote.Vote{StoryID: *id, UserID: userID, Created: time.Now()}
	}
	v.Choice = voteStr
	err := s.v.Save(ctx, nil, logger, v)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save vote for story")
	}
	return fe, "Vote recorded", currStory.PublicWebPath(fe.Estimate.Slug), nil
}

func estimateStoryRemove(
	ctx context.Context, fe *FullEstimate, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullEstimate, string, string, error) {
	id, _ := frm.GetUUID("storyID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	for _, v := range fe.Votes.GetByStoryIDs(*id) {
		err := s.v.Delete(ctx, nil, v.StoryID, v.UserID, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "unable to delete vote")
		}
	}
	curr := fe.Stories.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
	}
	err := s.st.Delete(ctx, nil, *id, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete story")
	}
	return fe, "Story deleted", fe.Estimate.PublicWebPath(), nil
}

func estimateMemberUpdate(ctx context.Context, fe *FullEstimate, frm util.ValueMap, s *Service, logger util.Logger) (*FullEstimate, string, string, error) {
	if fe.Self == nil {
		return nil, "", "", errors.New("you are not a member of this estimate")
	}
	if fe.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this estimate")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.em.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Member updated", fe.Estimate.PublicWebPath(), nil
}

func estimateMemberRemove(ctx context.Context, fe *FullEstimate, frm util.ValueMap, s *Service, logger util.Logger) (*FullEstimate, string, string, error) {
	if fe.Self == nil {
		return nil, "", "", errors.New("you are not a member of this estimate")
	}
	if fe.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this estimate")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fe.Members.Get(fe.Estimate.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this estimate", userID.String())
	}
	err := s.em.Delete(ctx, nil, curr.EstimateID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Member removed", fe.Estimate.PublicWebPath(), nil
}

func estimateUpdateSelf(ctx context.Context, fe *FullEstimate, frm util.ValueMap, s *Service, logger util.Logger) (*FullEstimate, string, string, error) {
	if fe.Self == nil {
		return nil, "", "", errors.New("you are not a member of this estimate")
	}
	choice := frm.GetStringOpt("choice")
	name := frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fe.Self.Name = name
	err := s.em.Update(ctx, nil, fe.Self, logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	return fe, "Profile edited", fe.Estimate.PublicWebPath(), nil
}
