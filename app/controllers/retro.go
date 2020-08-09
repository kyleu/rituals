package controllers

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/app/retro"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Retro(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))
		teams := app.Team(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Sprint(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcRetro.PluralTitle
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcRetro.Plural)}
		return npncontroller.T(templates.RetroList(sessions, teams, sprints, auths, params.Get(util.SvcRetro.Key, ctx.Logger), ctx, w))
	})
}

func RetroNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		categoriesString := r.Form.Get(npncore.Plural(npncore.KeyCategory))
		categories := npndatabase.StringToArray(categoriesString)
		if len(categories) == 0 {
			categories = retro.DefaultCategories
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcRetro, r.Form, ctx.App.User())

		sess, err := app.Retro(ctx.App).New(sf.Title, ctx.Profile.UserID, sf.MemberName, categories, sf.TeamID, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "error creating retro session")
		}

		_, err = app.Retro(ctx.App).Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return npncontroller.EResp(err, "error setting permissions for new session")
		}

		err = app.Socket(ctx.App).SendContentUpdate(util.SvcTeam, sf.TeamID)
		if err != nil {
			return npncontroller.EResp(err, "cannot send content update")
		}
		err = app.Socket(ctx.App).SendContentUpdate(util.SvcSprint, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcRetro.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func RetroWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Retro(ctx.App).GetBySlug(key)
		if sess == nil {
			msg := "can't load retro [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcRetro.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcRetro.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &web.PermissionParams{Svc: util.SvcRetro, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := web.CheckPerms(ctx, app.Retro(ctx.App).Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return web.PermErrorTemplate(util.SvcRetro, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return npncontroller.T(templates.RetroWorkspace(sess, auths, ctx, w))
	})
}

func RetroExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) web.ExportParams {
		sess := app.Retro(ctx.App).GetBySlug(key)
		if sess == nil {
			return web.ExportParams{}
		}
		return web.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcRetro.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Retro(ctx.App).Data.Permissions,
		}
	}
	web.ExportAct(util.SvcRetro, f, w, r)
}
