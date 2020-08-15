package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnconnection"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroSessionJoined struct {
	Profile     *npnuser.Profile       `json:"profile"`
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

func onRetroConnect(s *npnconnection.Service, conn *npnconnection.Connection, retroID uuid.UUID) error {
	ch := npnconnection.Channel{Svc: util.SvcRetro.Key, ID: retroID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinRetroSession(s, conn, ch)
	return errors.Wrap(err, "error joining retro session")
}

func joinRetroSession(s *npnconnection.Service, conn *npnconnection.Connection, ch npnconnection.Channel) error {
	dataSvc := retros(s)
	if ch.Svc != util.SvcRetro.Key {
		return errors.New("retro cannot handle [" + ch.Svc + "] message")
	}

	sess := dataSvc.GetByID(ch.ID)
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
		Comments:    dataSvc.Data.GetComments(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Sprint:      getSprintOpt(s, sess.SprintID),
		Members:     dataSvc.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Feedback:    dataSvc.GetFeedback(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := npnconnection.NewMessage(util.SvcRetro.Key, ServerCmdSessionJoined, sj)
	return sendInitial(s, ch, conn, res.Entry, msg, sess.SprintID, res.SprintEntry)
}
