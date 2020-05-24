package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
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
	Profile     *util.Profile          `json:"profile"`
	Session     *sprint.Session        `json:"session"`
	Permissions permission.Permissions `json:"permissions"`
	Auths       auth.Displays          `json:"auths"`
	Team        *team.Session          `json:"team"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
}

func onSprintConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	sprintID, err := uuid.FromString(param)
	if err != nil {
		return util.IDError(util.SvcTeam.Key, param)
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

	auths, displays := s.auths.GetDisplayByUserID(userID, nil)
	perms, permErrors, err := s.check(conn.Profile.UserID, auths, sess.TeamID, nil, util.SvcSprint, ch.ID)
	if err != nil {
		return err
	}
	if len(permErrors) > 0 {
		return s.sendPermErrors(util.SvcSprint, ch, permErrors)
	}

	entry := s.sprints.Members.Register(ch.ID, userID)
	members := s.sprints.Members.GetByModelID(ch.ID, nil)

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

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
			Auths:       displays,
			Team:        getTeamOpt(s, sess.TeamID),
			Members:     members,
			Online:      s.GetOnline(ch),
			Estimates:   estimates,
			Standups:    standups,
			Retros:      retros,
		},
	}

	return s.sendInitial(ch, conn, entry, msg, nil, nil)
}
