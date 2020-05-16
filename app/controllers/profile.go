package controllers

import (
	"emperror.dev/errors"
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, 0)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load auth records for user [" + ctx.Profile.UserID.String() + "]"))
		}
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("profile"), "profile")
		return tmpl(templates.Profile(auths, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		ctx.Profile.Name = strings.TrimSpace(r.Form.Get("username"))
		if ctx.Profile.Name == "" {
			ctx.Profile.Name = "Guest"
		}
		ctx.Profile.Theme = util.ThemeFromString(r.Form.Get("theme"))
		ctx.Profile.NavColor = r.Form.Get("navbar-color")
		ctx.Profile.LinkColor = r.Form.Get("link-color")
		_, err := ctx.App.User.SaveProfile(ctx.Profile)
		if err != nil {
			return "", err
		}
		ctx.Session.AddFlash("success:Profile saved")
		saveSession(w, r, ctx)
		return ctx.Route("home"), nil
	})
}
