package transcript

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type RetroResponse struct {
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
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil {
			return app.Retro.List(nil), nil
		}
		sprintID := param.(uuid.UUID)
		sess := app.Retro.GetByID(sprintID)
		dataSvc := app.Retro.Data
		return RetroResponse{
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Comments:    dataSvc.GetComments(sprintID, nil),
			Members:     dataSvc.Members.GetByModelID(sprintID, nil),
			Feedback:    app.Retro.GetFeedback(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sprintID, nil),
		}, nil
	},
}
