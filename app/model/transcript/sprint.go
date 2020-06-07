package transcript

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
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

type SprintResponse struct {
	Session     *sprint.Session        `json:"session"`
	Team        *team.Session          `json:"team"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
	Permissions permission.Permissions `json:"permissions"`
}

var Sprint = Transcript{
	Key:         util.SvcSprint.Key,
	Title:       util.SvcSprint.Title,
	Description: util.SvcSprint.Description,
	Resolve: func(app *config.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		if len(slug) == 0 {
			return app.Sprint.List(nil), nil
		}
		sess := app.Sprint.GetBySlug(slug)
		dataSvc := app.Sprint.Data
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		return SprintResponse{
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Estimates:   app.Estimate.GetBySprintID(sess.ID, nil),
			Standups:    app.Standup.GetBySprintID(sess.ID, nil),
			Retros:      app.Retro.GetBySprintID(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
