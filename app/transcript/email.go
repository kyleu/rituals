package transcript

import (
	"fmt"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"time"

	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/npn/npnservice/user"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
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
	return fmt.Sprintf("[%v] %v report", npncore.ToYMD(er.Date), npncore.AppName)
}

func (er *EmailResponse) Opener() string {
	msg := "Today there were %v users, %v auths, %v teams, %v sprints, %v estimates, %v standups, %v retros, and %v comments"
	return fmt.Sprintf(msg, len(er.Users), len(er.Auths), len(er.Teams), len(er.Sprints), len(er.Estimates), len(er.Standups), len(er.Retros), len(er.Comments))
}

var Email = Transcript{
	Key:         "email",
	Title:       "Email",
	Description: "Nightly email report",
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, param string) (interface{}, error) {
		if len(param) == 0 {
			n := time.Now()
			param = npncore.ToYMD(&n)
		}
		d, err := npncore.FromYMD(param)
		if err != nil {
			return nil, err
		}
		return EmailResponse{
			Date:      d,
			Users:     ai.User().GetByCreated(d, nil),
			Auths:     ai.Auth().GetByCreated(d, nil),
			Teams:     app.Team(ai).GetByCreated(d, nil),
			Sprints:   app.Sprint(ai).GetByCreated(d, nil),
			Estimates: app.Estimate(ai).GetByCreated(d, nil),
			Standups:  app.Standup(ai).GetByCreated(d, nil),
			Retros:    app.Retro(ai).GetByCreated(d, nil),
			Comments:  app.Comment(ai).GetByCreated(d, nil),
		}, nil
	},
}
