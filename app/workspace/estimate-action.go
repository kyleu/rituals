package workspace

import (
	"context"
	"github.com/google/uuid"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate/story"
	"github.com/kyleu/rituals/app/util"
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (s *Service) ActionEstimate(
	ctx context.Context, slug string, act string, frm util.ValueMap, userID uuid.UUID, logger util.Logger,
) (*FullEstimate, string, string, error) {
	w, err := s.LoadEstimate(ctx, slug, userID, nil, nil, logger)
	if err != nil {
		return nil, "", "", err
	}
	e := w.Estimate

	switch act {
	case "edit":
		tgt := e.Clone()
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
		return w, "Estimate saved", model.PublicWebPath(), nil
	case "story-add":
		title := strings.TrimSpace(frm.GetStringOpt("title"))
		if title == "" {
			return nil, "", "", errors.New("must provide [title]")
		}
		st := &story.Story{ID: util.UUID(), EstimateID: e.ID, Idx: len(w.Stories), UserID: userID, Title: title, Status: enum.SessionStatusNew, Created: time.Now()}
		err = s.st.Create(ctx, nil, logger, st)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "unable to save edited story")
		}
		return w, "Story added", e.PublicWebPath(), nil
	case "story-edit":
		id, _ := frm.GetUUID("storyID", false)
		if id == nil {
			return nil, "", "", errors.New("must provide [id]")
		}
		curr := w.Stories.Get(*id)
		if curr == nil {
			return nil, "", "", errors.Errorf("no story found with id [%s]", id.String())
		}
		title := strings.TrimSpace(frm.GetStringOpt("title"))
		if title == "" {
			return nil, "", "", errors.New("must provide [title]")
		}
		st := &story.Story{ID: *id, EstimateID: e.ID, Idx: curr.Idx, UserID: userID, Title: title, Status: curr.Status, Created: curr.Created}
		err = s.st.Update(ctx, nil, st, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "unable to save edited story")
		}
		return w, "Story saved", e.PublicWebPath(), nil
	case "self":
		if w.Self == nil {
			return nil, "", "", errors.New("you are not a member of this estimate")
		}
		choice := frm.GetStringOpt("choice")
		name := frm.GetStringOpt("name")
		if name == "" {
			return nil, "", "", errors.New("must provide [name]")
		}
		w.Self.Name = name
		err = s.em.Update(ctx, nil, w.Self, logger)
		if err != nil {
			return nil, "", "", errors.Wrap(err, "")
		}
		if choice == "global" {
			return nil, "", "", errors.New("can't change global name yet!")
		}
		return w, "Profile edited", e.PublicWebPath(), nil
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", act)
	}
}
