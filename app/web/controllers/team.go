package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		sessions := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))

		ctx.Title = util.PluralTitle(util.SvcTeam.Key)
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcTeam.Key+".list"), util.SvcTeam.Plural)
		return tmpl(templates.TeamList(sessions, auths, params.Get(util.SvcTeam.Key, ctx.Logger), ctx, w))
	})
}

func TeamNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()

		sf, err := parseSessionForm(ctx.Profile.UserID, util.SvcTeam, r.Form, ctx.App.User)
		if err != nil {
			return eresp(err, "cannot parse form")
		}

		sess, err := ctx.App.Team.New(sf.Title, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error creating team session")
		}

		_, err = ctx.App.Team.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error setting permissions for new session")
		}

		return ctx.Route(util.SvcTeam.Key, util.KeyKey, sess.Slug), nil
	})
}

func TeamWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess := ctx.App.Team.GetBySlug(key)
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcTeam.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcTeam, ModelID: sess.ID, Slug: key, Title: sess.Title}
		auths, permErrors, bc := check(&ctx, ctx.App.Team.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcTeam, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title

		return tmpl(templates.TeamWorkspace(sess, auths, ctx, w))
	})
}
