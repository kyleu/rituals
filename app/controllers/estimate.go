package controllers

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/app/estimate"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Estimate(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		teams := app.Team(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Sprint(ctx.App).GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcEstimate.PluralTitle
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcEstimate.Plural)}
		return npncontroller.T(templates.EstimateList(sessions, teams, sprints, auths, params.Get(util.SvcEstimate.Key, ctx.Logger), ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		choicesString := r.Form.Get(npncore.Plural(npncore.KeyChoice))
		choices := npndatabase.StringToArray(choicesString)
		if len(choices) == 0 {
			choices = estimate.DefaultChoices
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcEstimate, r.Form, ctx.App.User())

		sess, err := app.Estimate(ctx.App).New(sf.Title, ctx.Profile.UserID, sf.MemberName, choices, sf.TeamID, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "error creating estimate session")
		}

		_, err = app.Estimate(ctx.App).Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
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

		return ctx.Route(util.SvcEstimate.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Estimate(ctx.App).GetBySlug(key)
		if sess == nil {
			msg := "can't load estimate [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcEstimate.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcEstimate.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &web.PermissionParams{Svc: util.SvcEstimate, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := web.CheckPerms(ctx, app.Estimate(ctx.App).Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return web.PermErrorTemplate(util.SvcEstimate, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return npncontroller.T(templates.EstimateWorkspace(sess, auths, ctx, w))
	})
}

func EstimateExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) web.ExportParams {
		sess := app.Estimate(ctx.App).GetBySlug(key)
		if sess == nil {
			return web.ExportParams{}
		}
		return web.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcEstimate.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Estimate(ctx.App).Data.Permissions,
		}
	}
	web.ExportAct(util.SvcEstimate, f, w, r)
}
