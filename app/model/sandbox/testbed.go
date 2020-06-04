package sandbox

import (
	"github.com/kyleu/rituals.dev/app/model/email"
	"github.com/kyleu/rituals.dev/app/web"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx web.RequestContext) (interface{}, error) {
		es := email.NewService(ctx.App.Database, ctx.Logger)
		err := es.SendNightlyEmail(ctx.App, ctx.Profile.UserID, nil, false)
		if err != nil {
			return nil, err
		}
		return "Testbed!", nil
	},
}
