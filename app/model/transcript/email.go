package transcript

import (
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/model/comment"
	"github.com/kyleu/rituals.dev/app/model/estimate"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/model/sprint"
	"github.com/kyleu/rituals.dev/app/model/standup"
	"github.com/kyleu/rituals.dev/app/model/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EmailResponse struct {
	Date      *time.Time
	Teams     team.Sessions     `json:"teams"`
	Sprints   sprint.Sessions   `json:"sprints"`
	Estimates estimate.Sessions `json:"estimates"`
	Standups  standup.Sessions  `json:"standups"`
	Retros    retro.Sessions    `json:"retros"`
	Comments  comment.Comments  `json:"comments"`
}

func (er *EmailResponse) Subject() string {
	return fmt.Sprintf("[%v] %v report", util.ToYMD(er.Date), util.AppName)
}

func (er *EmailResponse) Opener() string {
	msg := "Today there were %v teams, %v sprints, %v estimates, %v standups, %v retros, and %v comments"
	return fmt.Sprintf(msg, len(er.Teams), len(er.Sprints), len(er.Estimates), len(er.Standups), len(er.Retros), len(er.Comments))
}

var Email = Transcript{
	Key:         "email",
	Title:       "Email",
	Description: "Nightly email report",
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param interface{}, format string) (interface{}, error) {
		if param == nil || len(param.(string)) == 0 {
			n := time.Now()
			param = util.ToYMD(&n)
		}
		d, err := util.FromYMD(param.(string))
		if err != nil {
			return nil, err
		}
		return EmailResponse{
			Date:      d,
			Teams:     app.Team.GetByCreated(d, nil),
			Sprints:   app.Sprint.GetByCreated(d, nil),
			Estimates: app.Estimate.GetByCreated(d, nil),
			Standups:  app.Standup.GetByCreated(d, nil),
			Retros:    app.Retro.GetByCreated(d, nil),
			Comments:  app.Comment.GetByCreated(d, nil),
		}, nil
	},
}
