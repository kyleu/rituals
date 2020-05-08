package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/util"
)

func onEstimateConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	estimateID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error reading channel id"))
	}
	ch := channel{Svc: util.SvcEstimate, ID: estimateID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinEstimateSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining estimate session"))
}

func joinEstimateSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcEstimate {
		return errors.WithStack(errors.New("estimate cannot handle [" + ch.Svc + "] message"))
	}

	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}
	if est == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing initial message"))
		}
		return nil
	}

	entry, _, err := s.estimates.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining estimate as member"))
	}

	members, err := s.estimates.Members.GetByModelID(ch.ID)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	stories, err := s.estimates.GetStories(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding stories"))
	}

	votes, err := s.estimates.GetEstimateVotes(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding votes"))
	}

	msg := Message{
		Svc: util.SvcEstimate,
		Cmd: util.ServerCmdSessionJoined,
		Param: EstimateSessionJoined{
			Profile: &conn.Profile,
			Session: est,
			Members: members,
			Online:  online,
			Stories: stories,
			Votes:   votes,
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
