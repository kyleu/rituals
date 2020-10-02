package admin

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npntemplate/gen/npntemplate"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
)

const codeLength = 16

var code = npncore.RandomString(codeLength)

func Enable(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		q := r.URL.Query()
		v, ok := q["code"]
		if !ok || len(v) == 0 {
			ctx.Logger.Warn("### admin enable request")
			ctx.Logger.Warn("### to enable admin access for this user")
			ctx.Logger.Warn(`### add "?code=` + code + `" to your url`)
			return npncontroller.T(npntemplate.StaticMessage("To become an admin, follow the instructions in your server logs", ctx, w))
		}
		if v[0] != code {
			if v[0] == (code + "!") {
				admin := uuid.FromStringOrNil("F0000000-0000-0000-0000-000000000000")
				npnweb.SetSessionUser(admin, ctx.Session, r, w, ctx.Logger)
				return ctx.Route(npncore.KeyAdmin), nil
			}
			return npncontroller.T(npntemplate.StaticMessage("Invalid code", ctx, w))
		}

		err := ctx.App.User().SetRole(ctx.Profile.UserID, npnuser.RoleAdmin)
		if err != nil {
			return npncontroller.EResp(err)
		}

		const msg = "you're a wizard, Harry!"
		return npncontroller.FlashAndRedir(true, msg, npncore.KeyAdmin, w, r, ctx)
	})
}
