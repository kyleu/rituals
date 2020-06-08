package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type StandupResponse struct {
	Svc         util.Service           `json:"-"`
	Session     *standup.Session       `json:"session"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Reports     standup.Reports        `json:"reports"`
	Permissions permission.Permissions `json:"permissions"`
}

var Standup = Transcript{
	Key:         util.SvcStandup.Key,
	Title:       util.SvcStandup.Title,
	Description: util.SvcStandup.Description,
	Resolve: func(app *config.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		if len(slug) == 0 {
			return app.Standup.List(nil), nil
		}
		sess := app.Standup.GetBySlug(slug)
		dataSvc := app.Standup.Data
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		return StandupResponse{
			Svc:         util.SvcStandup,
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Sprint:      app.Sprint.GetByIDPointer(sess.SprintID),
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Reports:     app.Standup.GetReports(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
