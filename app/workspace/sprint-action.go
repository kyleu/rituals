package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionSprint(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullSprint, string, string, error) {
	fs, err := s.LoadSprint(ctx, slug, userID, "", nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case action.ActUpdate:
		return sprintUpdate(ctx, fs, frm, slug, s, logger)
	case action.ActMemberUpdate:
		return sprintMemberUpdate(ctx, fs, frm, s, logger)
	case action.ActMemberRemove:
		return sprintMemberRemove(ctx, fs, frm, s, logger)
	case action.ActMemberSelf:
		return sprintUpdateSelf(ctx, fs, frm, s, logger)
	case action.ActComment:
		return sprintComment(ctx, fs, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func sprintUpdate(ctx context.Context, fs *FullSprint, frm util.ValueMap, slug string, s *Service, logger util.Logger) (*FullSprint, string, string, error) {
	tgt := fs.Sprint.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.r.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.rh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	tgt.StartDate, _ = frm.GetTime("startDate", false)
	tgt.EndDate, _ = frm.GetTime("endDate", false)
	tgt.TeamID, _ = frm.GetUUID(util.KeyTeam, true)
	model, err := s.SaveSprint(ctx, tgt, fs.Self.UserID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	fs.Sprint = model
	return fs, "Sprint saved", model.PublicWebPath(), nil
}

func sprintMemberUpdate(ctx context.Context, fs *FullSprint, frm util.ValueMap, s *Service, logger util.Logger) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	if fs.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this sprint")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.sm.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, curr, &fs.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Member updated", fs.Sprint.PublicWebPath(), nil
}

func sprintMemberRemove(ctx context.Context, fs *FullSprint, frm util.ValueMap, s *Service, logger util.Logger) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	if fs.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this sprint")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := fs.Members.Get(fs.Sprint.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this sprint", userID.String())
	}
	err := s.sm.Delete(ctx, nil, curr.SprintID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberRemove, userID, &fs.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Member removed", fs.Sprint.PublicWebPath(), nil
}

func sprintUpdateSelf(ctx context.Context, fs *FullSprint, frm util.ValueMap, s *Service, logger util.Logger) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	choice := frm.GetStringOpt("choice")
	name := frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	fs.Self.Name = name
	err := s.sm.Update(ctx, nil, fs.Self, logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": fs.Self.UserID, "name": name}
	err = s.send(enum.ModelServiceSprint, fs.Sprint.ID, action.ActMemberUpdate, arg, &fs.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Profile edited", fs.Sprint.PublicWebPath(), nil
}

func sprintComment(ctx context.Context, fs *FullSprint, frm util.ValueMap, s *Service, logger util.Logger) (*FullSprint, string, string, error) {
	if fs.Self == nil {
		return nil, "", "", errors.New("you are not a member of this sprint")
	}
	c, u, err := commentFromForm(frm, fs.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceSprint:
		if c.ModelID != fs.Sprint.ID {
			return nil, "", "", errors.New("this comment refers to a different sprint")
		}
	case enum.ModelServiceEstimate:
		if curr := fs.Estimates.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an estimate that isn't part of this sprint")
		}
	case enum.ModelServiceStandup:
		if curr := fs.Standups.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an standup that isn't part of this sprint")
		}
	case enum.ModelServiceRetro:
		if curr := fs.Retros.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an retro that isn't part of this sprint")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = s.c.Save(ctx, nil, logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceSprint, c.ModelID, action.ActComment, c, &fs.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fs, "Comment added", fs.Sprint.PublicWebPath() + u, nil
}
