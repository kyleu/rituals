package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/rituals.dev/app/comment"
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
	Profile     *npnuser.Profile       `json:"profile"`
	Session     *team.Session          `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Sprints     sprint.Sessions        `json:"sprints"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onTeamConnect(s *Service, conn *connection, teamID uuid.UUID) error {
	ch := Channel{Svc: util.SvcTeam, ID: teamID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinTeamSession(s, conn, ch)
	return errors.Wrap(err, "error joining team session")
}

func joinTeamSession(s *Service, conn *connection, ch Channel) error {
	dataSvc := s.teams
	if ch.Svc != util.SvcTeam {
		return errors.New("team cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := dataSvc.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, nil, nil, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := TeamSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    dataSvc.Data.GetComments(ch.ID, nil),
		Members:     dataSvc.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Sprints:     s.sprints.GetByTeamID(ch.ID, nil),
		Estimates:   s.estimates.GetByTeamID(ch.ID, nil),
		Standups:    s.standups.GetByTeamID(ch.ID, nil),
		Retros:      s.retros.GetByTeamID(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcTeam, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, nil, nil)
}
