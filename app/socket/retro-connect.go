package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *retro.Session         `json:"session"`
	Permissions permission.Permissions `json:"permissions"`
	Auths       auth.Displays          `json:"auths"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Feedback    []*retro.Feedback      `json:"feedback"`
}

func onRetroConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	retroID, err := uuid.FromString(param)
	if err != nil {
		return util.IDError(util.SvcRetro.Key, param)
	}
	ch := channel{Svc: util.SvcRetro.Key, ID: retroID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinRetroSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining retro session"))
}

func joinRetroSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcRetro.Key {
		return errors.WithStack(errors.New("retro cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.retros.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding retro session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcRetro.Key, Cmd: ServerCmdError, Param: "invalid session"})
		return errors.WithStack(errors.Wrap(err, "error writing error message"))
	}

	auths, displays := s.auths.GetDisplayByUserID(userID, nil)
	perms, permErrors, err := s.check(conn.Profile.UserID, auths, sess.TeamID, sess.SprintID, util.SvcRetro, ch.ID)
	if err != nil {
		return err
	}
	if len(permErrors) > 0 {
		return s.sendPermErrors(util.SvcRetro, ch, permErrors)
	}

	entry := s.retros.Members.Register(ch.ID, userID)
	sprintEntry := s.sprints.Members.RegisterRef(sess.SprintID, userID)
	members := s.retros.Members.GetByModelID(ch.ID, nil)

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

	feedback, err := s.retros.GetFeedback(ch.ID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding feedback for retro"))
	}

	msg := Message{
		Svc: util.SvcRetro.Key,
		Cmd: ServerCmdSessionJoined,
		Param: RetroSessionJoined{
			Profile:     &conn.Profile,
			Session:     sess,
			Permissions: perms,
			Auths:       displays,
			Team:        getTeamOpt(s, sess.TeamID),
			Sprint:      getSprintOpt(s, sess.SprintID),
			Members:     members,
			Online:      s.GetOnline(ch),
			Feedback:    feedback,
		},
	}

	return s.sendInitial(ch, conn, entry, msg, sess.SprintID, sprintEntry)
}
