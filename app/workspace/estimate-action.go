package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) ActionEstimate(
	ctx context.Context, slug string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullEstimate, string, string, error) {
	e, err := s.LoadEstimate(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act := frm.GetStringOpt("action"); act {
	case "edit":
		tgt := e.Estimate.Clone()
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
		return e, "Estimate saved", model.PublicWebPath(), nil
	case "story-add":
		return nil, "", "", errors.Errorf("TODO: %s", act)
	case "story-edit":
		return nil, "", "", errors.Errorf("TODO: %s", act)
	case "self":
		if e.Self == nil {
			return nil, "", "", errors.New("you are not a member of this estimate")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		e.Self.Name = name
		err = s.em.Update(ctx, nil, e.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return e, "Profile edited", e.Estimate.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
