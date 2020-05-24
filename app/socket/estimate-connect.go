package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateSessionJoined struct {
	Profile     *util.Profile          `json:"profile"`
	Session     *estimate.Session      `json:"session"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Stories     []*estimate.Story      `json:"stories"`
	Votes       []*estimate.Vote       `json:"votes"`
}

func onEstimateConnect(s *Service, conn *connection, userID uuid.UUID, param string) error {
	estimateID, err := uuid.FromString(param)
	if err != nil {
		return util.IDError(util.SvcEstimate.Key, param)
	}
	ch := channel{Svc: util.SvcEstimate.Key, ID: estimateID}
	err = s.Join(conn.ID, ch)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error joining channel"))
	}
	err = joinEstimateSession(s, conn, userID, ch)
	return errors.WithStack(errors.Wrap(err, "error joining estimate session"))
}

func joinEstimateSession(s *Service, conn *connection, userID uuid.UUID, ch channel) error {
	if ch.Svc != util.SvcEstimate.Key {
		return errors.WithStack(errors.New("estimate cannot handle [" + ch.Svc + "] message"))
	}

	sess, err := s.estimates.GetByID(ch.ID)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding estimate session"))
	}
	if sess == nil {
		err = s.WriteMessage(conn.ID, &Message{Svc: util.SvcEstimate.Key, Cmd: ServerCmdError, Param: "invalid session"})
		if err != nil {
			return errors.WithStack(errors.Wrap(err, "error writing error message"))
		}
		return nil
	}

	auths, displays := s.auths.GetDisplayByUserID(userID, nil)
	perms, permErrors, err := s.check(conn.Profile.UserID, auths, sess.TeamID, sess.SprintID, util.SvcEstimate, ch.ID)
	if err != nil {
		return err
	}
	if len(permErrors) > 0 {
		return s.sendPermErrors(util.SvcEstimate, ch, permErrors)
	}

	entry := s.estimates.Members.Register(ch.ID, userID)
	sprintEntry := s.sprints.Members.RegisterRef(sess.SprintID, userID)
	members := s.estimates.Members.GetByModelID(ch.ID, nil)

	stories, err := s.estimates.GetStories(ch.ID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding stories"))
	}

	votes, err := s.estimates.GetEstimateVotes(ch.ID, nil)
	if err != nil {
		return errors.WithStack(errors.Wrap(err, "error finding votes"))
	}

	conn.Svc = ch.Svc
	conn.ModelID = &ch.ID
	s.actions.Post(ch.Svc, ch.ID, userID, action.ActConnect, nil, "")

	msg := Message{
		Svc: util.SvcEstimate.Key,
		Cmd: ServerCmdSessionJoined,
		Param: EstimateSessionJoined{
			Profile:     &conn.Profile,
			Session:     sess,
			Auths:       displays,
			Permissions: perms,
			Team:        getTeamOpt(s, sess.TeamID),
			Sprint:      getSprintOpt(s, sess.SprintID),
			Members:     members,
			Online:      s.GetOnline(ch),
			Stories:     stories,
			Votes:       votes,
		},
	}

	return s.sendInitial(ch, conn, entry, msg, sess.SprintID, sprintEntry)
}
