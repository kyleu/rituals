package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *standup.Session       `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Reports     standup.Reports        `json:"reports"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onStandupConnect(s *Service, conn *Connection, standupID uuid.UUID) error {
	ch := Channel{Svc: util.SvcStandup, ID: standupID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinStandupSession(s, conn, ch)
	return errors.Wrap(err, "error joining standup session")
}

func joinStandupSession(s *Service, conn *Connection, ch Channel) error {
	if ch.Svc != util.SvcStandup {
		return errors.New("standup cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := s.standups.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, sess.TeamID, sess.SprintID, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := StandupSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    s.standups.Data.Comments.GetByModelID(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     s.standups.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Reports:     s.standups.GetReports(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcStandup, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
