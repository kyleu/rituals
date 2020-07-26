package transcript

import (
	"fmt"
	"time"

	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/user"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type EmailResponse struct {
	Date      *time.Time
	Users     user.SystemUsers  `json:"users"`
	Auths     auth.Records      `json:"records"`
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
	msg := "Today there were %v users, %v auths, %v teams, %v sprints, %v estimates, %v standups, %v retros, and %v comments"
	return fmt.Sprintf(msg, len(er.Users), len(er.Auths), len(er.Teams), len(er.Sprints), len(er.Estimates), len(er.Standups), len(er.Retros), len(er.Comments))
}

var Email = Transcript{
	Key:         "email",
	Title:       "Email",
	Description: "Nightly email report",
	Resolve: func(app *config.AppInfo, userID uuid.UUID, param string) (interface{}, error) {
		if len(param) == 0 {
			n := time.Now()
			param = util.ToYMD(&n)
		}
		d, err := util.FromYMD(param)
		if err != nil {
			return nil, err
		}
		return EmailResponse{
			Date:      d,
			Users:     app.User.GetByCreated(d, nil),
			Auths:     app.Auth.GetByCreated(d, nil),
			Teams:     app.Team.GetByCreated(d, nil),
			Sprints:   app.Sprint.GetByCreated(d, nil),
			Estimates: app.Estimate.GetByCreated(d, nil),
			Standups:  app.Standup.GetByCreated(d, nil),
			Retros:    app.Retro.GetByCreated(d, nil),
			Comments:  app.Comment.GetByCreated(d, nil),
		}, nil
	},
}
