package transcript

import (
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
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil {
			return app.Sprint.List(nil), nil
		}
		sprintID := param.(uuid.UUID)
		sess := app.Sprint.GetByID(sprintID)
		dataSvc := app.Sprint.Data
		return SprintResponse{
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Comments:    dataSvc.GetComments(sprintID, nil),
			Members:     dataSvc.Members.GetByModelID(sprintID, nil),
			Estimates:   app.Estimate.GetBySprintID(sprintID, nil),
			Standups:    app.Standup.GetBySprintID(sprintID, nil),
			Retros:      app.Retro.GetBySprintID(sprintID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sprintID, nil),
		}, nil
	},
}
