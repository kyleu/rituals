package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateSessionJoined struct {
	Profile     *npnuser.Profile          `json:"profile"`
	Session     *estimate.Session      `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Stories     estimate.Stories       `json:"stories"`
	Votes       estimate.Votes         `json:"votes"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onEstimateConnect(s *Service, conn *connection, estimateID uuid.UUID) error {
	ch := Channel{Svc: util.SvcEstimate, ID: estimateID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinEstimateSession(s, conn, ch)
	return errors.Wrap(err, "error joining estimate session")
}

func joinEstimateSession(s *Service, conn *connection, ch Channel) error {
	dataSvc := s.estimates
	if ch.Svc != util.SvcEstimate {
		return errors.New("estimate cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := dataSvc.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, sess.TeamID, sess.SprintID, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := EstimateSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    dataSvc.Data.GetComments(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     dataSvc.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Stories:     dataSvc.GetStories(ch.ID, nil),
		Votes:       dataSvc.GetEstimateVotes(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}

	msg := NewMessage(util.SvcEstimate, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
