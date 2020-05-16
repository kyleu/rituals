package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupSessionJoined struct {
	Profile *util.Profile    `json:"profile"`
	Session *standup.Session `json:"session"`
	Members []*member.Entry   `json:"members"`
	Online  []uuid.UUID      `json:"online"`
	Reports []*standup.Report `json:"reports"`
}

func onStandupMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case util.ClientCmdConnect:
		err = onStandupConnect(s, conn, userID, param.(string))
	case util.ClientCmdUpdateSession:
		err = onStandupSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdAddReport:
		err = onAddReport(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdUpdateReport:
		err = onEditReport(s, *conn.Channel, userID, param.(map[string]interface{}))
	case util.ClientCmdRemoveReport:
		err = onRemoveReport(s, *conn.Channel, userID, param.(string))
	default:
		err = errors.New("unhandled standup command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling standup message"))
}

func onStandupSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	s.logger.Debug(fmt.Sprintf("saving standup session [%s]", title))

	err := s.standups.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating standup session"))
	}

	err = sendStandupSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending standup session"))
}

func sendStandupSessionUpdate(s *Service, ch channel) error {
	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding standup session [" + ch.ID.String() + "]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load standup session [" + ch.ID.String() + "]"))
	}
	msg := Message{Svc: util.SvcStandup.Key, Cmd: util.ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending standup session"))
}
