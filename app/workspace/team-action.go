package workspace

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionTeam(
	ctx context.Context, slug string, act action.Act, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullTeam, string, string, error) {
	p := NewLoadParams(ctx, slug, userID, "", nil, nil, logger)
	ft, err := s.LoadTeam(p)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case action.ActUpdate:
		return teamUpdate(ctx, ft, frm, slug, s, logger)
	case action.ActMemberUpdate:
		return teamMemberUpdate(ctx, ft, frm, s, logger)
	case action.ActMemberRemove:
		return teamMemberRemove(ctx, ft, frm, s, logger)
	case action.ActMemberSelf:
		return teamUpdateSelf(ctx, ft, frm, s, logger)
	case action.ActComment:
		return teamComment(ctx, ft, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func teamUpdate(ctx context.Context, ft *FullTeam, frm util.ValueMap, slug string, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	tgt := ft.Team.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.r.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.rh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	model, err := s.SaveTeam(ctx, tgt, ft.Self.UserID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	ft.Team = model
	err = s.send(enum.ModelServiceTeam, ft.Team.ID, action.ActUpdate, model, &ft.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Team saved", model.PublicWebPath(), nil
}

func teamMemberUpdate(ctx context.Context, ft *FullTeam, frm util.ValueMap, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	if ft.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this team")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.tm.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, curr, &ft.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Member updated", ft.Team.PublicWebPath(), nil
}

func teamMemberRemove(ctx context.Context, ft *FullTeam, frm util.ValueMap, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	if ft.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this team")
	}
	userID, _ := frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	err := s.tm.Delete(ctx, nil, curr.TeamID, curr.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberRemove, userID, &ft.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Member removed", ft.Team.PublicWebPath(), nil
}

func teamUpdateSelf(ctx context.Context, ft *FullTeam, frm util.ValueMap, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	choice := frm.GetStringOpt("choice")
	name := frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	ft.Self.Name = name
	err := s.tm.Update(ctx, nil, ft.Self, logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == "global" {
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": ft.Self.UserID, "name": name}
	err = s.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, arg, &ft.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Profile edited", ft.Team.PublicWebPath(), nil
}

func teamComment(ctx context.Context, ft *FullTeam, frm util.ValueMap, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	c, u, err := commentFromForm(frm, ft.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceTeam:
		if c.ModelID != ft.Team.ID {
			return nil, "", "", errors.New("this comment refers to a different team")
		}
	case enum.ModelServiceSprint:
		if curr := ft.Sprints.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an sprint that isn't part of this team")
		}
	case enum.ModelServiceEstimate:
		if curr := ft.Estimates.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an estimate that isn't part of this team")
		}
	case enum.ModelServiceStandup:
		if curr := ft.Standups.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an standup that isn't part of this team")
		}
	case enum.ModelServiceRetro:
		if curr := ft.Retros.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an retro that isn't part of this team")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = s.c.Save(ctx, nil, logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = s.send(enum.ModelServiceTeam, c.ModelID, action.ActComment, c, &ft.Self.UserID, logger)
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Comment added", ft.Team.PublicWebPath() + u, nil
}
