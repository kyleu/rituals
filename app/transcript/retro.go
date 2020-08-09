package transcript

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroResponse struct {
	Svc         util.Service           `json:"-"`
	Session     *retro.Session         `json:"session"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Feedback    retro.Feedbacks        `json:"feedback"`
	Permissions permission.Permissions `json:"permissions"`
}

var Retro = Transcript{
	Key:         util.SvcRetro.Key,
	Title:       util.SvcRetro.Title,
	Description: util.SvcRetro.Description,
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, slug string) (interface{}, error) {
		if len(slug) == 0 {
			return app.Retro(ai).List(nil), nil
		}
		sess := app.Retro(ai).GetBySlug(slug)
		if sess == nil {
			return nil, errors.New("no session available matching [" + slug + "]")
		}
		dataSvc := app.Retro(ai).Data
		return RetroResponse{
			Svc:         util.SvcRetro,
			Session:     sess,
			Team:        app.Team(ai).GetByIDPointer(sess.TeamID),
			Sprint:      app.Sprint(ai).GetByIDPointer(sess.SprintID),
			Comments:    dataSvc.GetComments(sess.ID, nil),
			Members:     dataSvc.Members.GetByModelID(sess.ID, nil),
			Feedback:    app.Retro(ai).GetFeedback(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sess.ID, nil),
		}, nil
	},
}
