package socket

import (
	"fmt"
	"strings"

	"github.com/kyleu/rituals.dev/app/team"

	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/sprint"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroSessionJoined struct {
	Profile  *util.Profile     `json:"profile"`
	Session  *retro.Session    `json:"session"`
	Team     *team.Session     `json:"team"`
	Sprint   *sprint.Session   `json:"sprint"`
	Members  []*member.Entry   `json:"members"`
	Online   []uuid.UUID       `json:"online"`
	Feedback []*retro.Feedback `json:"feedback"`
}

func onRetroMessage(s *Service, conn *connection, userID uuid.UUID, cmd string, param interface{}) error {
	var err error
	switch cmd {
	case ClientCmdConnect:
		err = onRetroConnect(s, conn, userID, param.(string))
	case ClientCmdUpdateSession:
		err = onRetroSessionSave(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdAddFeedback:
		err = onAddFeedback(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdUpdateFeedback:
		err = onEditFeedback(s, *conn.Channel, userID, param.(map[string]interface{}))
	case ClientCmdRemoveFeedback:
		err = onRemoveFeedback(s, *conn.Channel, userID, param.(string))
	default:
		err = errors.New("unhandled retro command [" + cmd + "]")
	}
	return errors.WithStack(errors.Wrap(err, "error handling retro message"))
}

func onRetroSessionSave(s *Service, ch channel, userID uuid.UUID, param map[string]interface{}) error {
	title := util.ServiceTitle(param["title"].(string))
	categoriesString, ok := param["categories"].(string)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("cannot parse [%v] as string", param["categories"])))
	}
	categories := query.StringToArray(categoriesString)
	if len(categories) == 0 {
		categories = retro.DefaultCategories
	}

	sprintID := getUUIDPointer(param, "sprintID")
	teamID := getUUIDPointer(param, "teamID")

	s.logger.Debug(fmt.Sprintf("saving retro session [%s] with categories [%s] and sprint [%s]", title, strings.Join(categories, ", "), sprintID))

	curr, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error loading retro session ["+ch.ID.String()+"] for update"))
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	err = s.retros.UpdateSession(ch.ID, title, categories, teamID, sprintID, userID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error updating retro session"))
	}

	err = sendRetroSessionUpdate(s, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error sending retro session"))
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending team for updated retro session"))
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error sending sprint for updated retro session"))
		}
	}

	return nil
}

func sendRetroSessionUpdate(s *Service, ch channel) error {
	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding retro session ["+ch.ID.String()+"]"))
	}
	if sess == nil {
		return errors.WithStack(errors.Wrap(err, "cannot load retro session ["+ch.ID.String()+"]"))
	}

	msg := Message{Svc: util.SvcRetro.Key, Cmd: ServerCmdSessionUpdate, Param: sess}
	err = s.WriteChannel(ch, &msg)
	return errors.WithStack(errors.Wrap(err, "error sending retro session"))
}
