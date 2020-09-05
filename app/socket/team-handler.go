package socket

import (
	"fmt"
	"github.com/kyleu/npn/npnservice/auth"

	"github.com/kyleu/npn/npnconnection"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

func onTeamSessionSave(s *npnconnection.Service, a *auth.Service, ch npnconnection.Channel, userID uuid.UUID, param teamSessionSaveParams) error {
	dataSvc := teams(s)
	title := util.ServiceTitle(util.SvcTeam, param.Title)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no team available with id [" + ch.ID.String() + "]")
	}

	sr := checkPerms(s, a, userID, nil, nil, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	msg := "saving team session [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title))

	err := dataSvc.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.Wrap(err, "error updating team session")
	}

	if title != curr.Title {
		slug, err := dataSvc.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating team slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendTeamSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending team session")
	}

	err = updatePerms(s, ch, userID, nil, nil, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendTeamUpdate(s *npnconnection.Service, ch npnconnection.Channel, curr *uuid.UUID, tm *team.Session) error {
	err := s.WriteChannel(ch, npnconnection.NewMessage(ch.Svc, ServerCmdTeamUpdate, tm))
	if err != nil {
		return errors.Wrap(err, "error writing team update message")
	}

	err = SendContentUpdate(s, util.SvcTeam.Key, curr)
	if err != nil {
		return err
	}
	if tm != nil {
		err = SendContentUpdate(s, util.SvcTeam.Key, &tm.ID)
	}
	return err
}

func sendTeamSessionUpdate(s *npnconnection.Service, ch npnconnection.Channel) error {
	sess := teams(s).GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load team session [" + ch.ID.String() + "]")
	}
	err := s.WriteChannel(ch, npnconnection.NewMessage(util.SvcTeam.Key, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending team session")
}

func getTeamOpt(s *npnconnection.Service, teamID *uuid.UUID) *team.Session {
	if teamID == nil {
		return nil
	}
	return teams(s).GetByID(*teamID)
}
