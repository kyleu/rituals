package socket

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/util"
)

type SprintSessionJoined struct {
	Profile   *util.Profile       `json:"profile"`
	Session   *sprint.Session     `json:"session"`
	Members   []*member.Entry     `json:"members"`
	Online    []uuid.UUID         `json:"online"`
	Estimates []*estimate.Session `json:"estimates"`
	Standups  []*standup.Session  `json:"standups"`
	Retros    []*retro.Session    `json:"retros"`
}

func onSprintMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case ClientCmdConnect:
		err = onSprintConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onSprintSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	default:
		err = errors.New("unhandled sprint command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling sprint message"))
}

func onSprintSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	s.logger.Debug(fmt.Sprintf("saving sprint session [%s]", title))

	err := s.sprints.UpdateSession(ch.ID, title, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating sprint session"))
	}

	err = sendSprintSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending sprint session"))
}

func sendSprintSessionUpdate(s *Service, ch channel) error {
	sess, err := s.sprints.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding sprint session ["+ch.ID.String()+"]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load sprint session ["+ch.ID.String()+"]"))
	}
	msg := Message{Svc: util.SvcSprint.Key, Cmd: ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending sprint session"))
}
