package sandbox

import (
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/app/email"
	"github.com/kyleu/rituals.dev/app/transcript"
)

var Nightly = Sandbox{
	Key:         "nightly",
	Title:       "Nightly",
	Description: "View the nightly email report, optionally sending it",
	Resolve: func(ctx *npnweb.RequestContext) (string, interface{}, error) {
		es := email.NewService(app.Database(ctx.App), ctx.Logger)
		now := time.Now()
		ymd := npncore.ToYMD(&now)
		html, rsp, err := es.GetNightlyEmail(ymd, &transcript.Context{
			UserID: ctx.Profile.UserID,
			App:    ctx.App,
			Logger: ctx.App.Logger(),
			Routes: ctx.Routes,
		})
		if err != nil {
			return "", nil, err
		}
		return html, rsp, nil
	},
}
