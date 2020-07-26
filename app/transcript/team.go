package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/config"
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
	Resolve: func(app *config.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		if len(slug) == 0 {
			return app.Team.List(nil), nil
		}
		sess := app.Team.GetBySlug(slug)
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		dataSvc := app.Team.Data
		return TeamResponse{
			Svc:         util.SvcTeam,
			Session:     sess,
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Sprints:     app.Sprint.GetByTeamID(sess.ID, nil),
			Estimates:   app.Estimate.GetByTeamID(sess.ID, nil),
			Standups:    app.Standup.GetByTeamID(sess.ID, nil),
			Retros:      app.Retro.GetByTeamID(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
