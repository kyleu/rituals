package controllers

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Team(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = npncore.PluralTitle(util.SvcTeam.Key)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcTeam.Plural)}
		return npncontroller.T(templates.TeamList(sessions, auths, params.Get(util.SvcTeam.Key, ctx.Logger), ctx, w))
	})
}

func TeamNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcTeam, r.Form, ctx.App.User())

		sess, err := app.Team(ctx.App).New(sf.Title, ctx.Profile.UserID, sf.MemberName)
		if err != nil {
			return npncontroller.EResp(err, "error creating team session")
		}

		_, err = app.Team(ctx.App).Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return npncontroller.EResp(err, "error setting permissions for new session")
		}

		return ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func TeamWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Team(ctx.App).GetBySlug(key)
		if sess == nil {
			msg := "can't load team [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcTeam.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &web.PermissionParams{Svc: util.SvcTeam, ModelID: sess.ID, Slug: key, Title: sess.Title}
		auths, permErrors, bc := web.CheckPerms(ctx, app.Team(ctx.App).Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return web.PermErrorTemplate(util.SvcTeam, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title

		return npncontroller.T(templates.TeamWorkspace(sess, auths, ctx, w))
	})
}

func TeamExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) web.ExportParams {
		sess := app.Team(ctx.App).GetBySlug(key)
		if sess == nil {
			return web.ExportParams{}
		}
		return web.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Team(ctx.App).Data.Permissions,
		}
	}
	web.ExportAct(util.SvcTeam, f, w, r)
}
