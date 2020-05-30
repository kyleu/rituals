package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
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

func onEstimateConnect(s *Service, conn *Connection, estimateID uuid.UUID) error {
	ch := Channel{Svc: util.SvcEstimate, ID: estimateID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinEstimateSession(s, conn, ch)
	return errors.Wrap(err, "error joining estimate session")
}

func joinEstimateSession(s *Service, conn *Connection, ch Channel) error {
	if ch.Svc != util.SvcEstimate {
		return errors.New("estimate cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := s.estimates.GetByID(ch.ID)
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
		Comments:    s.estimates.Data.Comments.GetByModelID(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     s.estimates.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Stories:     s.estimates.GetStories(ch.ID, nil),
		Votes:       s.estimates.GetEstimateVotes(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}

	msg := NewMessage(util.SvcEstimate, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
