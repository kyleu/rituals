package socket

import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals.dev/app/team"

	"github.com/kyleu/rituals.dev/app/sprint"

	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateSessionJoined struct {
	Profile *util.Profile     `json:"profile"`
	Session *estimate.Session `json:"session"`
	Team    *team.Session     `json:"team"`
	Sprint  *sprint.Session   `json:"sprint"`
	Members []*member.Entry   `json:"members"`
	Online  []uuid.UUID       `json:"online"`
	Stories []*estimate.Story `json:"stories"`
	Votes   []*estimate.Vote  `json:"votes"`
}

func onEstimateMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case ClientCmdConnect:
		p, ok := param.(string)
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as string"))
		}
		err = onEstimateConnect(s, conn, userID, p)
	case ClientCmdUpdateSession:
		p, ok := param.(map[string]interface{})
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onEstimateSessionSave(s, *conn.Channel, userID, p)
	case ClientCmdAddStory:
		p, ok := param.(map[string]interface{})
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onAddStory(s, *conn.Channel, userID, p)
	case ClientCmdUpdateStory:
		p, ok := param.(map[string]interface{})
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onUpdateStory(s, *conn.Channel, userID, p)
	case ClientCmdRemoveStory:
		p, ok := param.(string)
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as string"))
		}
		err = onRemoveStory(s, *conn.Channel, userID, p)
	case ClientCmdSetStoryStatus:
		p, ok := param.(map[string]interface{})
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onSetStoryStatus(s, *conn.Channel, userID, p)
	case ClientCmdSubmitVote:
		p, ok := param.(map[string]interface{})
		if !ok {
			return errors.WithStack(errors.New("cannot read parameter as map[string]interface{}"))
		}
		err = onSubmitVote(s, *conn.Channel, userID, p)
	default:
		err = errors.New("unhandled estimate command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling estimate message"))
}

func onEstimateSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	titleString, ok := param["title"].(string)
	if !ok {
		return errors.WithStack(errors.New("cannot read choices as string"))
	}
	title := util.ServiceTitle(titleString)

	choicesString, ok := param["choices"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("cannot parse [%v] as string", param["choices"])))
	}
	choices := query.StringToArray(choicesString)
	if len(choices) == 0 {
		choices = estimate.DefaultChoices
	}

	sprintID := getUUIDPointer(param, "sprintID")
	teamID := getUUIDPointer(param, "teamID")

	msg := "saving estimate session [%s] with choices [%s], team [%s], and sprint [%s]"
	s.logger.Debug(fmt.Sprintf(msg, title, strings.Join(choices, ", "), teamID, sprintID))

	curr, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error loading estimate session ["+ch.ID.String()+"] for update"))
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.estimates.UpdateSession(ch.ID, title, choices, teamID, sprintID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating estimate session"))
	}

	err = sendEstimateSessionUpdate(s, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error sending estimate session"))
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending team for updated estimate session"))
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending sprint for updated estimate session"))
		}
	}

	return nil
}

func sendEstimateSessionUpdate(s *Service, ch channel) error {
	est, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding estimate session ["+ch.ID.String()+"]"))
	}
	if est == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load estimate session ["+ch.ID.String()+"]"))
	}

	msg := Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdSessionUpdate, Param: est}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending estimate session"))
}
