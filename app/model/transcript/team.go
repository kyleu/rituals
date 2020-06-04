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

type TeamResponse struct {
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
	Resolve:     func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil {
			return app.Team.List(nil), nil
		}
		teamID := param.(uuid.UUID)
		dataSvc := app.Sprint.Data
		return TeamResponse{
			Session:     app.Team.GetByID(teamID),
			Comments:    dataSvc.GetComments(teamID, nil),
			Members:     dataSvc.Members.GetByModelID(teamID, nil),
			Sprints:     app.Sprint.GetByTeamID(teamID, nil),
			Estimates:   app.Estimate.GetByTeamID(teamID, nil),
			Standups:    app.Standup.GetByTeamID(teamID, nil),
			Retros:      app.Retro.GetByTeamID(teamID, nil),
			Permissions: dataSvc.Permissions.GetByModelID(teamID, nil),
		}, nil
	},
}
