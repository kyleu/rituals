package socket

import (
	"fmt"
	"strings"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateSessionJoined struct {
	Profile *util.Profile     `json:"profile"`
	Session *estimate.Session `json:"session"`
	Members []*member.Entry    `json:"members"`
	Online  []uuid.UUID       `json:"online"`
	Stories []*estimate.Story  `json:"stories"`
	Votes   []*estimate.Vote   `json:"votes"`
}

func onEstimateMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case util.ClientCmdConnect:
		p, ok := param.(string)
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as string"))
		}
		err = onEstimateConnect(s, conn, userID, p)
	case util.ClientCmdUpdateSession:
		p, ok := param.(map[string]interface{})
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onEstimateSessionSave(s, *conn.Channel, userID, p)
	case util.ClientCmdAddStory:
		p, ok := param.(map[string]interface{})
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onAddStory(s, *conn.Channel, userID, p)
	case util.ClientCmdUpdateStory:
		p, ok := param.(map[string]interface{})
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onUpdateStory(s, *conn.Channel, userID, p)
	case util.ClientCmdRemoveStory:
		p, ok := param.(string)
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as string"))
		}
		err = onRemoveStory(s, *conn.Channel, userID, p)
	case util.ClientCmdSetStoryStatus:
		p, ok := param.(map[string]interface{})
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onSetStoryStatus(s, *conn.Channel, userID, p)
	case util.ClientCmdSubmitVote:
		p, ok := param.(map[string]interface{})
		if(!ok) {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onSubmitVote(s, *conn.Channel, userID, p)
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}

func onEstimateSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	titleString, ok := param["choices"].(string)
	if(!ok) {
		return errors.WithStack(errors.New("cannot read choices as string"))
	}
	title := util.ServiceTitle(titleString)
	choicesString, ok := param["choices"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("cannot parse [%v] as string", param["choices"])))
	}
	choices := util.StringToArray(choicesString)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}
	s.logger.Debug(fmt.Sprintf("saving estimate session [%s] with choices [%s]", title, strings.Join(choices, ", ")))

	err := s.estimates.UpdateSession(ch.ID, title, choices, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating estimate session"))
	}

	err = sendEstimateSessionUpdate(s, ch)
	return errors.WithStack(errors.Wrap(err, "error sending estimate session"))
}

func sendEstimateSessionUpdate(s *Service, ch channel) error {
	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding estimate session [" + ch.ID.String() + "]"))
	}
	if est == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load estimate session [" + ch.ID.String() + "]"))
	}

	msg := Message{Svc: util.SvcEstimate.Key, Cmd: util.ServerCmdSessionUpdate, Param: est}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending estimate session"))
}
