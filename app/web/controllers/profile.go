package controllers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/model/auth"
	"github.com/kyleu/rituals.dev/app/web/form"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if !tempSecurityCheck(&ctx) {
			return tmpl(templates.StaticMessage("Coming soon!", ctx, w))
		}

		params := act.ParamSetFromRequest(r)
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		ctx.Title = "User Profile"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.KeyProfile), util.KeyProfile)
		return tmpl(templates.Profile(ctx.App.Auth.Enabled, auths, ctx, w))
	})
}

func ProfileSave(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		prof := &form.ProfileForm{}
		err := form.Decode(r, prof, ctx.Logger)
		if err != nil {
			return eresp(err, "")
		}

		if len(strings.TrimSpace(prof.Username)) == 0 {
			prof.Username = "Guest"
		}

		ctx.Profile.Name = strings.TrimSpace(prof.Username)
		ctx.Profile.Theme = util.ThemeFromString(prof.Theme)
		ctx.Profile.NavColor = prof.NavColor
		ctx.Profile.LinkColor = prof.LinkColor

		_ = ctx.App.User.GetByID(ctx.Profile.UserID, true)
		_, err = ctx.App.User.SaveProfile(ctx.Profile)
		if err != nil {
			return eresp(err, "")
		}
		ctx.Session.AddFlash("success:Profile saved")
		act.SaveSession(w, r, ctx)
		return ctx.Route("home"), nil
	})
}

func ProfilePic(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		id, err := act.IDFromParams(util.KeyID, mux.Vars(r))
		if err != nil {
			return eresp(err, "invalid id")
		}
		a := ctx.App.Auth.GetByID(*id)
		ctx.Profile.Picture = a.Picture
		_, err = ctx.App.User.SaveProfile(ctx.Profile)
		if err != nil {
			return eresp(err, "can't save profile")
		}

		return ctx.Route(util.KeyProfile), nil
	})
}
