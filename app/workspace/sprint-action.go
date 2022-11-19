package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) ActionSprint(
	ctx context.Context, slug string, act string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullSprint, string, string, error) {
	spr, err := s.LoadSprint(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case "edit":
		tgt := spr.Sprint.Clone()
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
		model, err := s.SaveSprint(ctx, tgt, userID, nil, logger)
		if err != nil {
			return nil, "", "", err
		}
		return spr, "Sprint saved", model.PublicWebPath(), nil
	case "self":
		if spr.Self == nil {
			return nil, "", "", errors.New("you are not a member of this sprint")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		spr.Self.Name = name
		err = s.sm.Update(ctx, nil, spr.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return spr, "Profile edited", spr.Sprint.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
