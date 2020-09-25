package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/rituals.dev/app/action"
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

func sendInitial(s *npnconnection.Service, ch npnconnection.Channel, conn *npnconnection.Connection, entry *member.Entry, msg *npnconnection.Message, sprintID *uuid.UUID, sprintEntry *member.Entry) error {
	err := s.WriteMessage(conn.ID, msg)
	if err != nil {
		return errors.Wrap(err, "error writing initial message")
	}

	if sprintEntry != nil {
		err = sendMemberUpdate(s, npnconnection.Channel{Svc: util.SvcSprint.Key, ID: *sprintID}, sprintEntry, conn.ID)
		if err != nil {
			return errors.Wrap(err, "error writing member update to sprint")
		}
	}

	err = sendMemberUpdate(s, *conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.Wrap(err, "error writing member update")
	}

	err = sendOnlineUpdate(s, ch, conn.ID, conn.Profile.UserID, true)
	return errors.Wrap(err, "error writing online update")
}

func getPerms(s *npnconnection.Service, a auth.Service, auths auth.Records, userID uuid.UUID, teamID *uuid.UUID, sprintID *uuid.UUID, ch npnconnection.Channel) (permission.Permissions, permission.Errors) {
	return check(s, a, userID, auths, teamID, sprintID, ch.Svc, ch.ID)
}

func getSessionResult(s *npnconnection.Service, a auth.Service, teamID *uuid.UUID, sprintID *uuid.UUID, ch npnconnection.Channel, conn *npnconnection.Connection) SessionResult {
	userID := conn.Profile.UserID
	auths, displays := a.GetDisplayByUserID(userID, nil)

	perms, permErrors := getPerms(s, a, auths, conn.Profile.UserID, teamID, sprintID, ch)
	if len(permErrors) > 0 {
		return SessionResult{Error: sendPermErrors(s, conn.ID, ch.Svc, permErrors)}
	}

	dataSvc := dataFor(s, ch.Svc)
	entry := dataSvc.Members.Register(ch.ID, userID, "", member.RoleMember)
	var sprintEntry *member.Entry
	if sprintID != nil {
		sprintEntry = ctx(s).sprints.Data.Members.RegisterRef(sprintID, userID, "", member.RoleMember)
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	ctx(s).actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil)

	return SessionResult{
		Auth:        displays,
		Perms:       perms,
		Entry:       entry,
		SprintEntry: sprintEntry,
	}
}

func sendPermErrors(s *npnconnection.Service, connID uuid.UUID, svc string, permErrors permission.Errors) error {
	if len(permErrors) > 0 {
		return s.WriteMessage(connID, npnconnection.NewMessage(svc, ServerCmdError, "insufficient permissions"))
	}
	return nil
}

func errorNoSession(s *npnconnection.Service, svc string, connID uuid.UUID, chID uuid.UUID) error {
	msg := npncore.IDErrorString(npncore.KeySession, chID.String())
	err := s.WriteMessage(connID, npnconnection.NewMessage(util.SvcSystem.Key, ServerCmdSessionRemove, msg))
	if err != nil {
		return errors.Wrap(err, "error writing error message")
	}
	return errors.New("no " + svc + " session available")
}
