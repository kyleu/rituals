package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onRetroConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	retroID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error reading channel id"))
	}
	ch := channel{Svc: util.SvcRetro, ID: retroID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinRetroSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining retro session"))
}

func joinRetroSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcRetro {
		return errors.WithStack(errors.New("retro cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcRetro, Cmd: util.ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing initial message"))
		}
		return nil
	}

	entry, _, err := s.retros.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining retro as member"))
	}

	members, err := s.retros.Members.GetByModelID(ch.ID)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcRetro,
		Cmd: util.ServerCmdSessionJoined,
		Param: RetroSessionJoined{
			Profile: &conn.Profile,
			Session: sess,
			Members: members,
			Online:  online,
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
