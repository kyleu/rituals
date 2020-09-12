package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Svc(ctx.App).Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
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

		sess, err := app.Svc(ctx.App).Team.New(sf.Title, ctx.Profile.UserID, sf.MemberName)
		if err != nil {
			return npncontroller.EResp(err, "error creating team session")
		}

		_, err = app.Svc(ctx.App).Team.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return npncontroller.EResp(err, "error setting permissions for new session")
		}

		return ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func TeamWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Svc(ctx.App).Team.GetBySlug(key)
		if sess == nil {
			msg := "can't load team [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcTeam.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &PermissionParams{Svc: util.SvcTeam, ModelID: sess.ID, Slug: key, Title: sess.Title}
		auths, permErrors, bc := CheckPerms(ctx, app.Svc(ctx.App).Team.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return PermErrorTemplate(util.SvcTeam, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title

		return npncontroller.T(templates.TeamWorkspace(sess, auths, ctx, w))
	})
}

func TeamExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) ExportParams {
		sess := app.Svc(ctx.App).Team.GetBySlug(key)
		if sess == nil {
			return ExportParams{}
		}
		return ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcTeam.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Svc(ctx.App).Team.Data.Permissions,
		}
	}
	ExportAct(util.SvcTeam, f, w, r)
}
