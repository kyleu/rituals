package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/member"
	"github.com/kyleu/rituals.dev/internal/app/util"
	"strings"
)

type EstimateSessionJoined struct {
	Profile *util.Profile     `json:"profile"`
	Session *estimate.Session `json:"session"`
	Members []member.Entry    `json:"members"`
	Online  []uuid.UUID       `json:"online"`
	Polls   []estimate.Poll   `json:"polls"`
	Votes   []estimate.Vote   `json:"votes"`
}

func onEstimateMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error = nil
	switch cmd {
	case util.ClientCmdConnect:
		err = onConnect(s, conn, userID, param.(string))
	case util.ClientCmdUpdateSession:
		err = onSessionSave(s, *conn.Channel, param.(map[string]interface{}))
	case util.ClientCmdAddPoll:
		err = onAddPoll(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdUpdatePoll:
		err = onUpdatePoll(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdSetPollStatus:
		err = onSetPollStatus(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdSubmitVote:
		err = onSubmitVote(s, *conn.Channel, userID, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}

func onSessionSave(s *Service, ch channel, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	choicesString := param["choices"].(string)
	choices := util.StringToArray(choicesString)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}
	s.logger.Debug(fmt.Sprintf("saving estimate session [%s] with choices [%s]", title, strings.Join(choices, ", ")))

	err := s.estimates.UpdateSession(ch.ID, title, choices)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating session"))
	}

	err = sendSessionUpdate(s, ch, nil)
	return errors.WithStack(errors.Wrap(err, "error sending estimate"))
}

func sendSessionUpdate(s *Service, ch channel, connID *uuid.UUID) error {
	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding session"))
	}
	msg := Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdSessionUpdate, Param: est}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending session"))
}

func onAddPoll(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	s.logger.Debug(fmt.Sprintf("adding poll [%s]", title))

	poll, err := s.estimates.NewPoll(ch.ID, title, userID)
	err = sendPollUpdate(s, ch, poll)
	return errors.WithStack(errors.Wrap(err, "error sending polls"))
}

func onUpdatePoll(s *Service, c channel, id uuid.UUID, m map[string]interface{}) error {
	s.logger.Debug("TODO: update poll")
	return nil
}

func onSetPollStatus(s *Service, c channel, id uuid.UUID, m map[string]interface{}) error {
	s.logger.Debug("TODO: set poll status")
	return nil
}

func onSubmitVote(s *Service, c channel, id uuid.UUID, m map[string]interface{}) error {
	s.logger.Debug("TODO: submit vote")
	return nil
}

func sendPollUpdate(s *Service, ch channel, poll *estimate.Poll) error {
	msg := Message{Svc: util.SvcEstimate, Cmd: util.ServerCmdPollUpdate, Param: poll}
	err := s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending poll update"))
}
