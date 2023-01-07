package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/retro/feedback"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionRetro(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullRetro, string, string, error) {
	fr, err := s.LoadRetro(ctx, slug, userID, "", nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case action.ActUpdate:
		return retroUpdate(ctx, fr, frm, slug, s, logger)
	case action.ActFeedbackAdd:
		return retroFeedbackAdd(ctx, fr, frm, s, logger)
	case action.ActFeedbackUpdate:
		return retroFeedbackUpdate(ctx, fr, frm, s, logger)
	case action.ActFeedbackRemove:
		return retroFeedbackRemove(ctx, fr, frm, s, logger)
	case action.ActMemberUpdate:
		return retroMemberUpdate(ctx, fr, frm, s, logger)
	case action.ActMemberRemove:
		return retroMemberRemove(ctx, fr, frm, s, logger)
	case action.ActMemberSelf:
		return retroUpdateSelf(ctx, fr, frm, s, logger)
	case action.ActComment:
		return retroComment(ctx, fr, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func retroUpdate(
	ctx context.Context, fr *FullRetro, frm util.ValueMap, slug string, s *Service, logger util.Logger,
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
	model, err := s.SaveRetro(ctx, tgt, fr.Self.UserID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	fr.Retro = model
	return fr, "Retro saved", model.PublicWebPath(), nil
}

func retroFeedbackAdd(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	category := frm.GetStringOpt("category")
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: util.UUID(), RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: time.Now(),
	}
	err := s.f.Create(ctx, nil, logger, f)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	return fr, "Feedback added", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackUpdate(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	id, _ := frm.GetUUID("feedbackID", false)
	if id == nil {
		return nil, "", "", errors.New("must provide [id]")
	}
	curr := fr.Feedbacks.Get(*id)
	if curr == nil {
		return nil, "", "", errors.Errorf("no feedback found with id [%s]", id.String())
	}
	category := frm.GetStringOpt("category")
	content := frm.GetStringOpt("content")
	html := util.ToHTML(content, true)
	f := &feedback.Feedback{
		ID: *id, RetroID: fr.Retro.ID, Category: category, UserID: fr.Self.UserID, Content: content, HTML: html, Created: curr.Created, Updated: util.TimeToday(),
	}
	err := s.f.Update(ctx, nil, f, logger)
	if err != nil {
		return nil, "", "", errors.Wrap(err, "unable to save edited feedback")
	}
	return fr, "Feedback saved", fr.Retro.PublicWebPath(), nil
}

func retroFeedbackRemove(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
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

func retroMemberUpdate(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
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
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.rm.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, curr, &fr.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Member updated", fr.Retro.PublicWebPath(), nil
}

func retroMemberRemove(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
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
		return nil, "", "", errors.Errorf("user [%s] is not a member of this retro", userID.String())
	}
	err := s.rm.Delete(ctx, nil, curr.RetroID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberRemove, userID, &fr.Self.UserID, logger)
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
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": fr.Self.UserID, "name": name}
	err = s.send(enum.ModelServiceRetro, fr.Retro.ID, action.ActMemberUpdate, arg, &fr.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Profile edited", fr.Retro.PublicWebPath(), nil
}

func retroComment(ctx context.Context, fr *FullRetro, frm util.ValueMap, s *Service, logger util.Logger) (*FullRetro, string, string, error) {
	if fr.Self == nil {
		return nil, "", "", errors.New("you are not a member of this retro")
	}
	c, u, err := commentFromForm(frm, fr.Self.UserID)
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
	err = s.c.Save(ctx, nil, logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceRetro, c.ModelID, action.ActComment, c, &fr.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fr, "Comment added", fr.Retro.PublicWebPath() + u, nil
}
