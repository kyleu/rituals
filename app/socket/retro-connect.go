package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *retro.Session         `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Feedback    retro.Feedbacks        `json:"feedback"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onRetroConnect(s *Service, conn *Connection, retroID uuid.UUID) error {
	ch := Channel{Svc: util.SvcRetro, ID: retroID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinRetroSession(s, conn, ch)
	return errors.Wrap(err, "error joining retro session")
}

func joinRetroSession(s *Service, conn *Connection, ch Channel) error {
	if ch.Svc != util.SvcRetro {
		return errors.New("retro cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := s.retros.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, sess.TeamID, sess.SprintID, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := RetroSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    s.retros.Data.Comments.GetByModelID(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     s.retros.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Feedback:    s.retros.GetFeedback(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcRetro, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
