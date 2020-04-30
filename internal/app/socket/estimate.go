package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
)

func onEstimateMessage(s *Service, connID uuid.UUID, userID uuid.UUID, cmd string, param interface{}) error {
	var err error = nil
	switch cmd {
	case "connect":
		estimateString := param.(string)
		estimateID, err := uuid.FromString(estimateString)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error reading channel id"))
		}
		err = s.Join(connID, estimateID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error joining channel"))
		}

		est, err := s.estimates.GetByID(estimateID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error finding session"))
		}
		err = s.WriteMessage(connID, Message{Svc: "estimate", Cmd: "detail", Param: est})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending estimate"))
		}

		err = s.estimates.Join(estimateID, userID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error joining estimate as member"))
		}

		members, err := s.estimates.GetMembers(estimateID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error finding members"))
		}
		err = s.WriteChannel(estimateID, Message{Svc: "estimate", Cmd: "members", Param: members})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending members"))
		}

		polls, err := s.estimates.GetPolls(estimateID)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error finding polls"))
		}
		err = s.WriteMessage(connID, Message{Svc: "estimate", Cmd: "polls", Param: polls})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending polls"))
		}
	default:
		err = errors.New("Unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}
