package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) ActionRetro(
	ctx context.Context, slug string, act string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullRetro, string, string, error) {
	r, err := s.LoadRetro(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act {
	case "edit":
		tgt := r.Retro.Clone()
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
		return r, "Retro saved", model.PublicWebPath(), nil
	case "self":
		if r.Self == nil {
			return nil, "", "", errors.New("you are not a member of this retro")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		r.Self.Name = name
		err = s.rm.Update(ctx, nil, r.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return r, "Profile edited", r.Retro.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
