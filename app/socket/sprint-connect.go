package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type SprintSessionJoined struct {
	Profile     *util.Profile            `json:"profile"`
	Session     *sprint.Session          `json:"session"`
	Permissions []*permission.Permission `json:"permissions"`
	Team        *team.Session            `json:"team"`
	Members     []*member.Entry          `json:"members"`
	Online      []uuid.UUID              `json:"online"`
	Estimates   []*estimate.Session      `json:"estimates"`
	Standups    []*standup.Session       `json:"standups"`
	Retros      []*retro.Session         `json:"retros"`
}

func onSprintConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	sprintID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.New("error reading channel id [" + param + "]"))
	}
	ch := channel{Svc: util.SvcSprint.Key, ID: sprintID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinSprintSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining sprint session"))
}

func joinSprintSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcSprint.Key {
		return errors.WithStack(errors.New("sprint cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.sprints.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding sprint session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcSprint.Key, Cmd: ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing sprint error message"))
		}
		return nil
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

	entry := s.sprints.Members.Register(ch.ID, userID)
	perms := s.sprints.Permissions.GetByModelID(ch.ID, nil)
	members := s.sprints.Members.GetByModelID(ch.ID, nil)

	estimates, err := s.estimates.GetBySprint(ch.ID, nil)
	if err != nil {
		return err
	}
	standups, err := s.standups.GetBySprint(ch.ID, nil)
	if err != nil {
		return err
	}
	retros, err := s.retros.GetBySprint(ch.ID, nil)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcSprint.Key,
		Cmd: ServerCmdSessionJoined,
		Param: SprintSessionJoined{
			Profile:     &conn.Profile,
			Session:     sess,
			Permissions: perms,
			Team:        getTeamOpt(s, sess.TeamID),
			Members:     members,
			Online:      s.GetOnline(ch),
			Estimates:   estimates,
			Standups:    standups,
			Retros:      retros,
		},
	}

	err = s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing initial sprint message"))
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing member update"))
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.WithStack(errors.Wrap(err, "error writing online update"))
}
