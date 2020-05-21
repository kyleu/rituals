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

type TeamSessionJoined struct {
	Profile     *util.Profile            `json:"profile"`
	Session     *team.Session            `json:"session"`
	Permissions []*permission.Permission `json:"permissions"`
	Members     []*member.Entry          `json:"members"`
	Online      []uuid.UUID              `json:"online"`
	Sprints     []*sprint.Session        `json:"sprints"`
	Estimates   []*estimate.Session      `json:"estimates"`
	Standups    []*standup.Session       `json:"standups"`
	Retros      []*retro.Session         `json:"retros"`
}

func onTeamConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	teamID, err := uuid.FromString(param)
	if err != nil {
		return errors.WithStack(errors.New("error reading channel id [" + param + "]"))
	}
	ch := channel{Svc: util.SvcTeam.Key, ID: teamID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinTeamSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining team session"))
}

func joinTeamSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcTeam.Key {
		return errors.WithStack(errors.New("team cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.teams.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding team session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcTeam.Key, Cmd: ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing team error message"))
		}
		return nil
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

	entry := s.teams.Members.Register(ch.ID, userID)
	perms := s.teams.Permissions.GetByModelID(ch.ID, nil)
	members := s.teams.Members.GetByModelID(ch.ID, nil)

	sprints, err := s.sprints.GetByTeamID(ch.ID, nil)
	if err != nil {
		return err
	}
	estimates, err := s.estimates.GetByTeamID(ch.ID, nil)
	if err != nil {
		return err
	}
	standups, err := s.standups.GetByTeamID(ch.ID, nil)
	if err != nil {
		return err
	}
	retros, err := s.retros.GetByTeamID(ch.ID, nil)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcTeam.Key,
		Cmd: ServerCmdSessionJoined,
		Param: TeamSessionJoined{
			Profile:     &conn.Profile,
			Session:     sess,
			Permissions: perms,
			Members:     members,
			Online:      s.GetOnline(ch),
			Sprints:     sprints,
			Estimates:   estimates,
			Standups:    standups,
			Retros:      retros,
		},
	}

	err = s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing initial team message"))
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error writing member update"))
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.WithStack(errors.Wrap(err, "error writing online update"))
}
