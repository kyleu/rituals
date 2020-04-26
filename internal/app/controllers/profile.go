package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/app/util"
	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("profile"), "profile")
		return templates.Profile(ctx, w)
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	redir(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		ctx.Profile.Name = r.Form.Get("username")
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
