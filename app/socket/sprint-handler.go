package socket

import (
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/util"
)

func onSprintSessionSave(s *Service, ch Channel, userID uuid.UUID, param sprintSessionSaveParams) error {
	title := util.ServiceTitle(util.SvcSprint, param.Title)

	curr := s.sprints.GetByID(ch.ID)

	teamID := util.GetUUIDFromString(param.TeamID)
	var startDate *time.Time
	var endDate *time.Time

	if len(param.StartDate) > 0 {
		d, e := util.FromYMD(param.StartDate)
		if e == nil {
			startDate = d
		}
	}
	if len(param.EndDate) > 0 {
		d, e := util.FromYMD(param.EndDate)
		if e == nil {
			endDate = d
		}
	}

	s.Logger.Debug(fmt.Sprintf("saving sprint session [%s] in team [%s]", title, teamID))

	teamChanged := differentPointerValues(curr.TeamID, teamID)

	err := s.sprints.UpdateSession(ch.ID, title, teamID, startDate, endDate, userID)
	if err != nil {
		return errors.Wrap(err, "error updating sprint session")
	}

	err = sendSprintSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending sprint session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated sprint session")
		}
	}

	err = s.updatePerms(ch, userID, teamID, nil, s.sprints.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendSprintUpdate(s *Service, ch Channel, curr *uuid.UUID, spr *sprint.Session) error {
	err := s.WriteChannel(ch, NewMessage(ch.Svc, ServerCmdSprintUpdate, spr))
	if err != nil {
		return errors.Wrap(err, "error writing sprint update message")
	}
	err = s.SendContentUpdate(util.SvcSprint, curr)
	if err != nil {
		return err
	}
	if spr != nil {
		err = s.SendContentUpdate(util.SvcSprint, &spr.ID)
	}
	return err
}

func sendSprintSessionUpdate(s *Service, ch Channel) error {
	sess := s.sprints.GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load sprint session ["+ch.ID.String()+"]")
	}
	err := s.WriteChannel(ch, NewMessage(util.SvcSprint, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending sprint session")
}

func getSprintOpt(s *Service, sprintID *uuid.UUID) *sprint.Session {
	if sprintID == nil {
		return nil
	}
	return s.sprints.GetByID(*sprintID)
}
