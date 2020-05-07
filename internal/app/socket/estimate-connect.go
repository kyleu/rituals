package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/util"
)

func onConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
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

	entry, _, err := s.estimates.Members.Register(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining estimate as member"))
	}

	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}

	members, err := s.estimates.Members.GetByModelID(ch.ID)
	if err != nil {
		return err
	}

	online, err := s.GetOnline(ch)
	if err != nil {
		return err
	}

	polls, err := s.estimates.GetPolls(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding polls"))
	}

	var votes []estimate.Vote

	msg := Message{
		Svc:   util.SvcEstimate,
		Cmd:   util.ServerCmdSessionJoined,
		Param: EstimateSessionJoined{
			Profile: &conn.Profile,
			Session: est,
			Members: members,
			Online: online,
			Polls: polls,
			Votes: votes,
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

	err =  s.sendOnlineUpdate(ch, conn.Profile.UserID, true)
	return errors.WithStack(errors.Wrap(err, "error writing online update"))
}
