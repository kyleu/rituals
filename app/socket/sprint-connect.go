package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type SprintSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *sprint.Session        `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onSprintConnect(s *Service, conn *Connection, sprintID uuid.UUID) error {
	ch := Channel{Svc: util.SvcSprint, ID: sprintID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinSprintSession(s, conn, ch)
	return errors.Wrap(err, "error joining sprint session")
}

func joinSprintSession(s *Service, conn *Connection, ch Channel) error {
	if ch.Svc != util.SvcSprint {
		return errors.New("sprint cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := s.sprints.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, sess.TeamID, nil, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := SprintSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    s.sprints.Data.Comments.GetByModelID(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Members:     s.sprints.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Estimates:   s.estimates.GetBySprint(ch.ID, nil),
		Standups:    s.standups.GetBySprint(ch.ID, nil),
		Retros:      s.retros.GetBySprint(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcSprint, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, nil, nil)
}
