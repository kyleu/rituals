package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
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

type TeamResponse struct {
	Svc         util.Service           `json:"-"`
	Session     *team.Session          `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Sprints     sprint.Sessions        `json:"sprints"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
	Permissions permission.Permissions `json:"permissions"`
}

var Team = Transcript{
	Key:         util.SvcTeam.Key,
	Title:       util.SvcTeam.Title,
	Description: util.SvcTeam.Description,
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		svc := app.Svc(ai)
		if len(slug) == 0 {
			return svc.Team.List(nil), nil
		}
		sess := svc.Team.GetBySlug(slug)
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		dataSvc := svc.Team.Data
		return TeamResponse{
			Svc:         util.SvcTeam,
			Session:     sess,
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Sprints:     svc.Sprint.GetByTeamID(sess.ID, nil),
			Estimates:   svc.Estimate.GetByTeamID(sess.ID, nil),
			Standups:    svc.Standup.GetByTeamID(sess.ID, nil),
			Retros:      svc.Retro.GetByTeamID(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
