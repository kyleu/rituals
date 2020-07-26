package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

type SessionResult struct {
	Auth        auth.Displays
	Perms       permission.Permissions
	Entry       *member.Entry
	SprintEntry *member.Entry
	Error       error
}

func (s *Service) sendInitial(ch Channel, conn *connection, entry *member.Entry, msg *Message, sprintID *uuid.UUID, sprintEntry *member.Entry) error {
	err := s.WriteMessage(conn.ID, msg)
	if err != nil {
		return errors.Wrap(err, "error writing initial message")
	}

	if sprintEntry != nil {
		err = s.sendMemberUpdate(Channel{Svc: util.SvcSprint, ID: *sprintID}, sprintEntry, conn.ID)
		if err != nil {
			return errors.Wrap(err, "error writing member update to sprint")
		}
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.Wrap(err, "error writing member update")
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.Wrap(err, "error writing online update")
}

func getPerms(s *Service, auths auth.Records, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID, ch Channel) (permission.Permissions, permission.Errors) {
	return s.check(userID, auths, teamID, sprintID, ch.Svc, ch.ID)
}

func getSessionResult(s *Service, teamID *uuid.UUID, sprintID *uuid.UUID, ch Channel, conn *connection) SessionResult {
	userID := conn.Profile.UserID
	auths, displays := s.auths.GetDisplayByUserID(userID, nil)

	perms, permErrors := getPerms(s, auths, conn.Profile.UserID, teamID, sprintID, ch)
	if len(permErrors) > 0 {
		return SessionResult{Error: s.sendPermErrors(conn.ID, ch.Svc, permErrors)}
	}

	dataSvc := dataFor(s, ch.Svc)
	entry := dataSvc.Members.Register(ch.ID, userID, "", member.RoleMember)
	var sprintEntry *member.Entry
	if sprintID != nil {
		sprintEntry = s.sprints.Data.Members.RegisterRef(sprintID, userID, "", member.RoleMember)
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil)

	return SessionResult{
		Auth:        displays,
		Perms:       perms,
		Entry:       entry,
		SprintEntry: sprintEntry,
	}
}

func (s *Service) sendPermErrors(connID uuid.UUID, svc util.Service, permErrors permission.Errors) error {
	if len(permErrors) > 0 {
		return s.WriteMessage(connID, NewMessage(svc, ServerCmdError, "insufficient permissions"))
	}
	return nil
}

func errorNoSession(s *Service, svc util.Service, connID uuid.UUID, chID uuid.UUID) error {
	msg := util.IDErrorString(util.KeySession, chID.String())
	err := s.WriteMessage(connID, NewMessage(util.SvcEstimate, ServerCmdSessionRemove, msg))
	if err != nil {
		return errors.Wrap(err, "error writing error message")
	}
	return errors.New("no " + svc.Key + " session available")
}
