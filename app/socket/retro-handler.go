package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onRetroSessionSave(s *Service, ch Channel, userID uuid.UUID, param retroSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcRetro, param.Title)

	categories := query.StringToArray(param.Categories)
	if len(categories) == 0 {
		categories = retro.DefaultCategories
	}

	sprintID := util.GetUUIDFromString(param.SprintID)
	teamID := util.GetUUIDFromString(param.TeamID)

	curr := s.retros.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no retro available with id [" + ch.ID.String() + "]")
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	msg := "saving retro session [%s] with categories [%s], sprint [%s] and team [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title, util.OxfordComma(categories, "and"), sprintID, teamID))

	err := s.retros.UpdateSession(ch.ID, title, categories, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating retro session")
	}

	if title != curr.Title {
		slug, err := s.retros.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating retro slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendRetroSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending retro session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated retro session")
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated retro session")
		}
	}

	err = s.updatePerms(ch, userID, teamID, sprintID, s.retros.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendRetroSessionUpdate(s *Service, ch Channel) error {
	sess := s.retros.GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load retro session [" + ch.ID.String() + "]")
	}

	err := s.WriteChannel(ch, NewMessage(util.SvcRetro, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending retro session")
}
