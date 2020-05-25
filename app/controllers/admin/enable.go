package admin

import (
	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/controllers/act"
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

var code = util.RandomString(16)

func Enable(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		q := r.URL.Query()
		v, ok := q["code"]
		if !ok || len(v) == 0 {
			ctx.Logger.Warn("### admin enable request")
			ctx.Logger.Warn("### to enable admin access for this user")
			ctx.Logger.Warn("### add \"?code=" + code + "\" to your url")
			return tmpl(templates.Todo("To become an admin, follow the instructions in your server logs", ctx, w))
		}
		if v[0] != code {
			return tmpl(templates.Todo("Invalid code", ctx, w))
		}

		err := ctx.App.User.SetRole(ctx.Profile.UserID, util.RoleAdmin)
		if err != nil {
			return "", errors.Wrap(err, "unable to set role")
		}

		ctx.Session.AddFlash("success:You're a wizard, Harry!")
		act.SaveSession(w, r, ctx)

		return ctx.Route(util.KeyAdmin), nil
	})
}
