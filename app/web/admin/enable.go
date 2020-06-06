package admin

import (
	"net/http"

	"github.com/gofrs/uuid"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

const codeLength = 16

var code = util.RandomString(codeLength)

func Enable(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		q := r.URL.Query()
		v, ok := q["code"]
		if !ok || len(v) == 0 {
			ctx.Logger.Warn("### admin enable request")
			ctx.Logger.Warn("### to enable admin access for this user")
			ctx.Logger.Warn("### add \"?code=" + code + "\" to your url")
			return tmpl(templates.StaticMessage("To become an admin, follow the instructions in your server logs", ctx, w))
		}
		if v[0] != code {
			if v[0] == (code + "!") {
				web.SetSessionUser(uuid.FromStringOrNil("F0000000-0000-0000-0000-000000000000"), ctx.Session, r, w, ctx.Logger)
				return ctx.Route(util.KeyAdmin), nil
			}
			return tmpl(templates.StaticMessage("Invalid code", ctx, w))
		}

		err := ctx.App.User.SetRole(ctx.Profile.UserID, util.RoleAdmin)
		if err != nil {
			return eresp(err, "")
		}

		ctx.Session.AddFlash("success:You're a wizard, Harry!")
		act.SaveSession(w, r, ctx)

		return ctx.Route(util.KeyAdmin), nil
	})
}
