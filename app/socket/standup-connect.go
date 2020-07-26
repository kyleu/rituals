package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/comment"
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
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Reports     standup.Reports        `json:"reports"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onStandupConnect(s *Service, conn *connection, standupID uuid.UUID) error {
	ch := Channel{Svc: util.SvcStandup, ID: standupID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinStandupSession(s, conn, ch)
	return errors.Wrap(err, "error joining standup session")
}

func joinStandupSession(s *Service, conn *connection, ch Channel) error {
	dataSvc := s.standups
	if ch.Svc != util.SvcStandup {
		return errors.New("standup cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := dataSvc.GetByID(ch.ID)
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
		Comments:    dataSvc.Data.GetComments(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     dataSvc.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Reports:     dataSvc.GetReports(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcStandup, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
