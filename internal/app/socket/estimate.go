package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"strings"
)

func onEstimateMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error = nil
	switch cmd {
	case "connect":
		err = onConnect(s, conn, userID, param.(string))
	case "session-save":
		err = onSessionSave(s, *conn.Channel, param.(map[string]interface {}))
	case "new-poll-save":
		err = onAddPoll(s, *conn.Channel, param.(map[string]interface {}))
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}

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
	err = joinEstimateSession(s, conn.ID, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining estimate session"))
}

func joinEstimateSession(s *Service, connID uuid.UUID, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcEstimate {
		return errors.WithStack(errors.New("estimate cannot handle [" + ch.Svc + "] message"))
	}

	err := sendSession(s, ch, &connID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error sending estimate"))
	}

	_, err = s.estimates.Join(ch.ID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining estimate as member"))
	}

	err = s.SendMembers(&s.estimates.Members, ch, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error sending members to channel"))
	}

	err = sendPolls(s, ch, &connID)

	return errors.WithStack(errors.Wrap(err, "error sending polls to socket"))
}

func sendPolls(s *Service, ch channel, connID *uuid.UUID) error {
	polls, err := s.estimates.GetPolls(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding polls"))
	}
	msg := Message{Svc: util.SvcEstimate, Cmd: "polls", Param: polls}
	if connID == nil {
		err = s.WriteChannel(ch, &msg)
	} else {
		err = s.WriteMessage(*connID, &msg)
	}
	return errors.WithStack(errors.Wrap(err, "error sending polls"))
}

func sendSession(s *Service, ch channel, connID *uuid.UUID) error {
	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}
	msg := Message{Svc: util.SvcEstimate, Cmd: "detail", Param: est}
	if connID == nil {
		err = s.WriteChannel(ch, &msg)
	} else {
		err = s.WriteMessage(*connID, &msg)
	}
	return errors.WithStack(errors.Wrap(err, "error sending session"))
}

func onSessionSave(s *Service, ch channel, param map[string]interface {}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	choicesString := param["choices"].(string)
	choices := util.StringToArray(choicesString)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}
	s.logger.Debug(fmt.Sprintf("Saving estimate session [%s] with choices [%s]", title, strings.Join(choices, ", ")))

	err := sendSession(s, ch, nil)
	return errors.WithStack(errors.Wrap(err, "error sending estimate"))
}

func onAddPoll(s *Service, ch channel, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	s.logger.Debug(fmt.Sprintf("Adding poll [%s]", title))

	err := sendPolls(s, ch, nil)
	return errors.WithStack(errors.Wrap(err, "error sending polls"))
}

