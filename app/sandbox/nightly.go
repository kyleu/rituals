package sandbox

import (
	"time"

	"github.com/kyleu/rituals.dev/app/email"
	"github.com/kyleu/rituals.dev/app/transcript"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var Nightly = Sandbox{
	Key:         "nightly",
	Title:       "Nightly",
	Description: "View the nightly email report, optionally sending it",
	Resolve: func(ctx *web.RequestContext) (string, interface{}, error) {
		es := email.NewService(ctx.App.Database, ctx.Logger)
		now := time.Now()
		ymd := util.ToYMD(&now)
		html, rsp, err := es.GetNightlyEmail(ymd, &transcript.Context{
			UserID: ctx.Profile.UserID,
			App:    ctx.App,
			Logger: ctx.App.Logger,
			Routes: ctx.Routes,
		})
		if err != nil {
			return "", nil, err
		}
		return html, rsp, nil
	},
}
