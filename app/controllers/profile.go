package controllers

import (
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if !securityCheck(&ctx) {
			return tmpl(templates.Todo("Coming soon!", ctx, w))
		}

		params := act.ParamSetFromRequest(r)
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return "", errors.Wrap(err, "cannot load auth records for user ["+ctx.Profile.UserID.String()+"]")
		}
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyProfile), util.KeyProfile)
		return tmpl(templates.Profile(ctx.App.Auth.Enabled, auths, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		ctx.Profile.Name = strings.TrimSpace(r.Form.Get("username"))
		if ctx.Profile.Name == "" {
			ctx.Profile.Name = "Guest"
		}
		ctx.Profile.Theme = util.ThemeFromString(r.Form.Get(util.KeyTheme))
		ctx.Profile.NavColor = r.Form.Get("navbar-color")
		ctx.Profile.LinkColor = r.Form.Get("link-color")
		_, err := ctx.App.User.SaveProfile(ctx.Profile)
		if err != nil {
			return "", err
		}
		ctx.Session.AddFlash("success:Profile saved")
		act.SaveSession(w, r, ctx)
		return ctx.Route("home"), nil
	})
}
