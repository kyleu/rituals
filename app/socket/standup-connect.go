package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onStandupConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	standupID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error reading channel id"))
	}
	ch := channel{Svc: util.SvcStandup, ID: standupID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinStandupSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining standup session"))
}

func joinStandupSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcStandup {
		return errors.WithStack(errors.New("standup cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcStandup, Cmd: util.ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing initial message"))
		}
		return nil
	}

	entry, _, err := s.standups.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining standup as member"))
	}

	members, err := s.standups.Members.GetByModelID(ch.ID)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	updates, err := s.standups.GetUpdates(ch.ID)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcStandup,
		Cmd: util.ServerCmdSessionJoined,
		Param: StandupSessionJoined{
			Profile: &conn.Profile,
			Session: sess,
			Members: members,
			Online:  online,
			Updates: updates,
		},
	}

	err = s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing initial message"))
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing member update"))
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.WithStack(errors.Wrap(err, "error writing online update"))
}
