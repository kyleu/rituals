package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
)

func (s *Service) ActionTeam(
	ctx context.Context, slug string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullTeam, string, string, error) {
	t, err := s.LoadTeam(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}

	switch act := frm.GetStringOpt("action"); act {
	case "edit":
		tgt := t.Team.Clone()
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
		return t, "Team saved", model.PublicWebPath(), nil
	case "self":
		if t.Self == nil {
			return nil, "", "", errors.New("you are not a member of this team")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		t.Self.Name = name
		err = s.tm.Update(ctx, nil, t.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return t, "Profile edited", t.Team.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
