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
	ft, err := s.LoadTeam(ctx, slug, userID, "", nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case action.ActUpdate:
		return teamUpdate(ctx, ft, userID, frm, slug, s, logger)
	case action.ActMemberUpdate:
		return teamMemberUpdate(ctx, ft, frm, s, logger)
	case action.ActMemberRemove:
		return teamMemberRemove(ctx, ft, frm, s, logger)
	case action.ActMemberSelf:
		return teamUpdateSelf(ctx, ft, frm, s, logger)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}

func teamUpdate(
	ctx context.Context, ft *FullTeam, userID uuid.UUID, frm util.ValueMap, slug string, s *Service, logger util.Logger,
) (*FullTeam, string, string, error) {
	tgt := ft.Team.Clone()
	tgt.Title = frm.GetStringOpt("title")
	tgt.Slug = frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tgt.Slug = s.r.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.rh, nil, logger)
	tgt.Icon = frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	model, err := s.SaveTeam(ctx, tgt, userID, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	ft.Team = model
	return ft, "Team saved", model.PublicWebPath(), nil
}

func teamMemberUpdate(ctx context.Context, fe *FullTeam, frm util.ValueMap, s *Service, logger util.Logger) (*FullTeam, string, string, error) {
	if fe.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	if fe.Self.Role != enum.MemberStatusOwner {
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
	curr := fe.Members.Get(fe.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := s.tm.Update(ctx, nil, curr, logger)
	if err != nil {
		return nil, "", "", err
	}
	return fe, "Member updated", fe.Team.PublicWebPath(), nil
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
	return ft, "Profile edited", ft.Team.PublicWebPath(), nil
}
