package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
	"time"
)

func (s *Service) ActionRetro(
	ctx context.Context, slug string, act string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullRetro, string, string, error) {
	fr, err := s.LoadRetro(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case "edit":
		return retroEdit(ctx, fr, userID, frm, slug, s, logger)
	case "feedback-add":
		return retroFeedbackAdd(ctx, fr, userID, frm, s, logger)
	case "feedback-edit":
		return retroFeedbackEdit(ctx, fr, userID, frm, s, logger)
	case "feedback-delete":
		return retroFeedbackDelete(ctx, fr, userID, frm, s, logger)
	case "member-edit":
		return retroMemberEdit(ctx, fr, frm, s, logger)
	case "member-leave":
		return retroMemberLeave(ctx, fr, frm, s, logger)
	case "self":
		return retroUpdateSelf(ctx, fr, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func retroEdit(
	ctx context.Context, fr *FullRetro, userID uuid.UUID, frm util.ValueMap, slug string, s *Service, logger util.Logger,
) (*FullRetro, string, string, error) {
	tgt := fr.Retro.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.r.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.rh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.Categories = util.StringSplitAndTrim(frm.GetStringOpt("categories"), ",")
	tgt.TeamID, _ = frm.GetUUID(util.KeyTeam, true)
	tgt.SprintID, _ = frm.GetUUID(util.KeySprint, true)
	model, err := s.SaveRetro(ctx, tgt, userID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	fr.Retro = model
	return fr, "Retro saved", model.PublicWebPath(), nil
}

func retroFeedbackAdd(
	ctx context.Context, fu *FullRetro, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullRetro, string, string, error) {
	category := frm.GetStringOpt("category")
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content)
	f := &feedback.Feedback{
		ID: util.UUID(), RetroID: fu.Retro.ID, Category: category, UserID: userID, Content: content, HTML: html, Created: time.Now(),
	}
	err := s.f.Create(ctx, nil, logger, f)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	return fu, "Feedback added", fu.Retro.PublicWebPath(), nil
}

func retroFeedbackEdit(
	ctx context.Context, fu *FullRetro, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullRetro, string, string, error) {
	id, _ := frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fu.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	category := frm.GetStringOpt("category")
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content)
	f := &feedback.Feedback{
		ID: *id, RetroID: fu.Retro.ID, Category: category, UserID: userID, Content: content, HTML: html, Created: curr.Created, Updated: util.TimeToday(),
	}
	err := s.f.Update(ctx, nil, f, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	return fu, "Feedback saved", fu.Retro.PublicWebPath(), nil
}

func retroFeedbackDelete(
	ctx context.Context, fr *FullRetro, userID uuid.UUID, frm util.ValueMap, s *Service, logger util.Logger,
) (*FullRetro, string, string, error) {
	id, _ := frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	err := s.f.Delete(ctx, nil, *id, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to delete feedback")
	}
	return fr, "Feedback deleted", fr.Retro.PublicWebPath(), nil
}

func retroMemberEdit(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	if fr.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this retro")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro")
	}
	curr.Role = enum.MemberStatus(role)
	err := s.rm.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Member updated", fr.Retro.PublicWebPath(), nil
}

func retroMemberLeave(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	if fr.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this retro")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fr.Members.Get(fr.Retro.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro")
	}
	err := s.rm.Delete(ctx, nil, curr.RetroID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Member removed", fr.Retro.PublicWebPath(), nil
}

func retroUpdateSelf(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	choice := frm.GetStringOpt("choice")
	name := frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fr.Self.Name = name
	err := s.rm.Update(ctx, nil, fr.Self, logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet!")
	}
	return fr, "Profile edited", fr.Retro.PublicWebPath(), nil
}
