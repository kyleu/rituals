package transcript

import (
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
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil {
			return app.Standup.List(nil), nil
		}
		sprintID := param.(uuid.UUID)
		sess := app.Standup.GetByID(sprintID)
		dataSvc := app.Standup.Data
		return StandupResponse{
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Comments:    dataSvc.GetComments(sprintID, nil),
			Members:     dataSvc.Members.GetByModelID(sprintID, nil),
			Reports:     app.Standup.GetReports(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sprintID, nil),
		}, nil
	},
}
