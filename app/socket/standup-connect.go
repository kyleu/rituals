package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *standup.Session       `json:"session"`
	Permissions permission.Permissions `json:"permissions"`
	Auths       auth.Displays          `json:"auths"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Reports     standup.Reports      `json:"reports"`
}

func onStandupConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	standupID, err := uuid.FromString(param)
	if err != nil {
		return util.IDError(util.SvcStandup.Key, param)
	}
	ch := channel{Svc: util.SvcStandup.Key, ID: standupID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinStandupSession(s, conn, userID, ch)
	return errors.Wrap(err, "error joining standup session")
}

func joinStandupSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcStandup.Key {
		return errors.New("standup cannot handle [" + ch.Svc + "] message")
	}

	sess, err := s.standups.GetByID(ch.ID)
	if err != nil {
		return errors.Wrap(err, "error finding standup session")
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcStandup.Key, Cmd: ServerCmdError, Param: util.IDErrorString(util.KeySession, "")})
		if err != nil {
			return errors.Wrap(err, "error writing standup error message")
		}
		return nil
	}

	auths, displays := s.auths.GetDisplayByUserID(userID, nil)

	perms, permErrors, err := s.check(conn.Profile.UserID, auths, sess.TeamID, sess.SprintID, util.SvcStandup, ch.ID)
	if err != nil {
		return err
	}
	if len(permErrors) > 0 {
		return s.sendPermErrors(util.SvcStandup, ch, permErrors)
	}

	entry := s.standups.Members.Register(ch.ID, userID)
	sprintEntry := s.sprints.Members.RegisterRef(sess.SprintID, userID)
	members := s.standups.Members.GetByModelID(ch.ID, nil)

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

	reports, err := s.standups.GetReports(ch.ID, nil)
	if err != nil {
		return err
	}

	msg := Message{
		Svc: util.SvcStandup.Key,
		Cmd: ServerCmdSessionJoined,
		Param: StandupSessionJoined{
			Profile:     &conn.Profile,
			Session:     sess,
			Permissions: perms,
			Auths:       displays,
			Team:        getTeamOpt(s, sess.TeamID),
			Sprint:      getSprintOpt(s, sess.SprintID),
			Members:     members,
			Online:      s.GetOnline(ch),
			Reports:     reports,
		},
	}

	return s.sendInitial(ch, conn, entry, msg, sess.SprintID, sprintEntry)
}
