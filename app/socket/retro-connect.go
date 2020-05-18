package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

func onRetroConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	retroID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.New("error reading channel id [" + param + "]"))
	}
	ch := channel{Svc: util.SvcRetro.Key, ID: retroID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinRetroSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining retro session"))
}

func joinRetroSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcRetro.Key {
		return errors.WithStack(errors.New("retro cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding retro session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcRetro.Key, Cmd: ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing error message"))
		}
		return nil
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, "connect", nil, "")

	entry, _, err := s.retros.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining retro as member"))
	}

	var sprintEntry *member.Entry
	if sess.SprintID != nil {
		sprintEntry, _, err = s.sprints.Members.Register(*sess.SprintID, userID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error joining sprint as member"))
		}
	}

	members, err := s.retros.Members.GetByModelID(ch.ID, nil)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	feedback, err := s.retros.GetFeedback(ch.ID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding feedback for retro"))
	}

	msg := Message{
		Svc: util.SvcRetro.Key,
		Cmd: ServerCmdSessionJoined,
		Param: RetroSessionJoined{
			Profile:  &conn.Profile,
			Session:  sess,
			Members:  members,
			Online:   online,
			Feedback: feedback,
		},
	}

	err = s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing initial retro message"))
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
