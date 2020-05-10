package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
	"strings"
)

type StandupSessionJoined struct {
	Profile *util.Profile    `json:"profile"`
	Session *standup.Session `json:"session"`
	Members []member.Entry   `json:"members"`
	Online  []uuid.UUID      `json:"online"`
	Updates []standup.Update `json:"updates"`
}

func onStandupMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error = nil
	switch cmd {
	case util.ClientCmdConnect:
		err = onStandupConnect(s, conn, userID, param.(string))
	case util.ClientCmdUpdateSession:
		err = onStandupSessionSave(s, *conn.Channel, param.(map[string]interface{}))
	case util.ClientCmdAddUpdate:
		err = onAddUpdate(s, *conn.Channel, userID, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling standup message"))
}

func onStandupSessionSave(s *Service, ch channel, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	s.logger.Debug(fmt.Sprintf("saving standup session [%s]", title))

	err := s.standups.UpdateSession(ch.ID, title)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating standup session"))
	}

	err = sendStandupSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending standup session"))
}

func sendStandupSessionUpdate(s *Service, ch channel) error {
	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding standup session"))
	}
	msg := Message{Svc: util.SvcStandup, Cmd: util.ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending standup session"))
}
