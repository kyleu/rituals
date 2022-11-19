package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) ActionStandup(
	ctx context.Context, slug string, act string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullStandup, string, string, error) {
	u, err := s.LoadStandup(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	switch act {
	case "edit":
		tgt := u.Standup.Clone()
		tgt.Title = frm.GetStringOpt("title")
		tgt.Slug = frm.GetStringOpt("slug")
		if tgt.Slug == "" {
			tgt.Slug = util.Slugify(tgt.Title)
		}
		tgt.Slug = s.u.Slugify(ctx, tgt.ID, tgt.Slug, slug, s.uh, nil, logger)
		tgt.Icon = frm.GetStringOpt("icon")
		tgt.Icon = tgt.IconSafe()
		tgt.TeamID, _ = frm.GetUUID(util.KeyTeam, true)
		tgt.SprintID, _ = frm.GetUUID(util.KeySprint, true)
		model, err := s.SaveStandup(ctx, tgt, userID, nil, logger)
		if err != nil {
			return nil, "", "", err
		}
		return u, "Standup saved", model.PublicWebPath(), nil
	case "self":
		if u.Self == nil {
			return nil, "", "", errors.New("you are not a member of this standup")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		u.Self.Name = name
		err = s.um.Update(ctx, nil, u.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return u, "Profile edited", u.Standup.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
