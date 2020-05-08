package socket

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
	"strings"
)

type RetroSessionJoined struct {
	Profile *util.Profile  `json:"profile"`
	Session *retro.Session `json:"session"`
	Members []member.Entry `json:"members"`
	Online  []uuid.UUID    `json:"online"`
}

func onRetroMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case util.ClientCmdConnect:
		err = onRetroConnect(s, conn, userID, param.(string))
	case util.ClientCmdUpdateSession:
		err = onRetroSessionSave(s, *conn.Channel, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling retro message"))
}

func onRetroSessionSave(s *Service, ch channel, param map[string]interface{}) error {
	title := strings.TrimSpace(param["title"].(string))
	if title == "" {
		title = "Untitled"
	}
	s.logger.Debug(fmt.Sprintf("saving retro session [%s]", title))

	err := s.retros.UpdateSession(ch.ID, title)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating retro session"))
	}

	err = sendRetroSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending retro session"))
}

func sendRetroSessionUpdate(s *Service, ch channel) error {
	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding retro session"))
	}
	msg := Message{Svc: util.SvcRetro, Cmd: util.ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending retro session"))
}
