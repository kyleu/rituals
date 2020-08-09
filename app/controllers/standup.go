package controllers

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Standup(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		teams := app.Team(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Sprint(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcStandup.Plural)}
		return npncontroller.T(templates.StandupList(sessions, teams, sprints, auths, params.Get(util.SvcStandup.Key, ctx.Logger), ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcStandup, r.Form, ctx.App.User())

		sess, err := app.Standup(ctx.App).New(sf.Title, ctx.Profile.UserID, sf.MemberName, sf.TeamID, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "error creating standup session")
		}

		_, err = app.Standup(ctx.App).Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
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

		return ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Standup(ctx.App).GetBySlug(key)
		if sess == nil {
			msg := "can't load standup [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcStandup.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &web.PermissionParams{Svc: util.SvcStandup, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := web.CheckPerms(ctx, app.Standup(ctx.App).Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return web.PermErrorTemplate(util.SvcStandup, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return npncontroller.T(templates.StandupWorkspace(sess, auths, ctx, w))
	})
}

func StandupExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) web.ExportParams {
		sess := app.Standup(ctx.App).GetBySlug(key)
		if sess == nil {
			return web.ExportParams{}
		}
		return web.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Standup(ctx.App).Data.Permissions,
		}
	}
	web.ExportAct(util.SvcStandup, f, w, r)
}
