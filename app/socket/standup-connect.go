package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func onStandupConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	standupID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.New("error reading channel id [" + param + "]"))
	}
	ch := channel{Svc: util.SvcStandup.Key, ID: standupID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinStandupSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining standup session"))
}

func joinStandupSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcStandup.Key {
		return errors.WithStack(errors.New("standup cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding standup session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcStandup.Key, Cmd: util.ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing standup error message"))
		}
		return nil
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, "connect", nil, "")

	entry, _, err := s.standups.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining standup as member"))
	}

	var sprintEntry *member.Entry
	if sess.SprintID != nil {
		sprintEntry, _, err = s.sprints.Members.Register(*sess.SprintID, userID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error joining sprint as member"))
		}
	}

	members, err := s.standups.Members.GetByModelID(ch.ID)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	reports, err := s.standups.GetReports(ch.ID)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcStandup.Key,
		Cmd: util.ServerCmdSessionJoined,
		Param: StandupSessionJoined{
			Profile: &conn.Profile,
			Session: sess,
			Members: members,
			Online:  online,
			Reports: reports,
		},
	}

	err = s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing initial standup message"))
	}

	if sprintEntry != nil {
		err = s.sendMemberUpdate(channel{Svc: util.SvcSprint.Key, ID: *sess.SprintID}, sprintEntry, conn.ID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing member update to sprint"))
		}
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing member update"))
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.WithStack(errors.Wrap(err, "error writing online update"))
}
