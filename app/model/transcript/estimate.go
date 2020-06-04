package transcript

import (
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/member"
	"github.com/kyleu/rituals.dev/app/model/permission"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EstimateResponse struct {
	Session     *estimate.Session      `json:"session"`
	Team        *team.Session          `json:"team"`
	Sprint      *sprint.Session        `json:"sprint"`
	Comments    comment.Comments       `json:"comments"`
	Members     member.Entries         `json:"members"`
	Stories     estimate.Stories       `json:"stories"`
	Votes       estimate.Votes         `json:"votes"`
	Permissions permission.Permissions `json:"permissions"`
}

var Estimate = Transcript{
	Key:         util.SvcEstimate.Key,
	Title:       util.SvcEstimate.Title,
	Description: util.SvcEstimate.Description,
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil {
			return app.Estimate.List(nil), nil
		}
		sprintID := param.(uuid.UUID)
		sess := app.Estimate.GetByID(sprintID)
		dataSvc := app.Estimate.Data
		return EstimateResponse{
			Session:     sess,
			Team:        app.Team.GetByIDPointer(sess.TeamID),
			Comments:    dataSvc.GetComments(sprintID, nil),
			Members:     dataSvc.Members.GetByModelID(sprintID, nil),
			Stories:     app.Estimate.GetStories(sess.ID, nil),
			Votes:       app.Estimate.GetEstimateVotes(sess.ID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(sprintID, nil),
		}, nil
	},
}
