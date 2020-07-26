package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		sessions := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		teams := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcEstimate.PluralTitle
		ctx.Breadcrumbs = web.Breadcrumbs{web.BreadcrumbSelf(util.SvcEstimate.Plural)}
		return act.T(templates.EstimateList(sessions, teams, sprints, auths, params.Get(util.SvcEstimate.Key, ctx.Logger), ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_ = r.ParseForm()

		choicesString := r.Form.Get(util.Plural(util.KeyChoice))
		choices := query.StringToArray(choicesString)
		if len(choices) == 0 {
			choices = estimate.DefaultChoices
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcEstimate, r.Form, ctx.App.User)

		sess, err := ctx.App.Estimate.New(sf.Title, ctx.Profile.UserID, sf.MemberName, choices, sf.TeamID, sf.SprintID)
		if err != nil {
			return act.EResp(err, "error creating estimate session")
		}

		_, err = ctx.App.Estimate.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return act.EResp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam, sf.TeamID)
		if err != nil {
			return act.EResp(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint, sf.SprintID)
		if err != nil {
			return act.EResp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcEstimate.Key, util.KeyKey, sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess := ctx.App.Estimate.GetBySlug(key)
		if sess == nil {
			msg := "can't load estimate [" + key + "]"
			return act.FlashAndRedir(false, msg, util.SvcEstimate.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcEstimate.Key, util.KeyKey, sess.Slug), nil
		}

		params := &act.PermissionParams{Svc: util.SvcEstimate, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := act.CheckPerms(ctx, ctx.App.Estimate.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return act.PermErrorTemplate(util.SvcEstimate, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return act.T(templates.EstimateWorkspace(sess, auths, ctx, w))
	})
}

func EstimateExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *web.RequestContext) act.ExportParams {
		sess := ctx.App.Estimate.GetBySlug(key)
		if sess == nil {
			return act.ExportParams{}
		}
		return act.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcEstimate.Key, util.KeyKey, sess.Slug),
			PermSvc: ctx.App.Estimate.Data.Permissions,
		}
	}
	act.ExportAct(util.SvcEstimate, f, w, r)
}
