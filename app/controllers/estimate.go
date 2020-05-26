package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/query"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "error retrieving estimates")
		}

		teams, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}
		sprints, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}

		ctx.Title = util.SvcEstimate.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key+".list"), util.SvcEstimate.Key)
		return tmpl(templates.EstimateList(sessions, teams, sprints, auths, params.Get(util.SvcEstimate.Key, ctx.Logger), ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()

		choicesString := r.Form.Get("choices")
		choices := query.StringToArray(choicesString)
		if len(choices) == 0 {
			choices = estimate.DefaultChoices
		}

		r, err := parseSessionForm(ctx.Profile.UserID, util.SvcEstimate, r.Form, ctx.App.User)
		if err != nil {
			return eresp(err, "cannot parse form")
		}

		sess, err := ctx.App.Estimate.New(r.Title, ctx.Profile.UserID, choices, r.TeamID, r.SprintID)
		if err != nil {
			return eresp(err, "error creating estimate session")
		}

		_, err = ctx.App.Estimate.Permissions.SetAll(sess.ID, r.Perms, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, r.TeamID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, r.SprintID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcEstimate.Key, util.KeyKey, sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Estimate.GetBySlug(key)
		if err != nil {
			return eresp(err, "cannot load estimate session")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcEstimate.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcEstimate, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := check(&ctx, ctx.App.Estimate.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcEstimate, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.EstimateWorkspace(sess, auths, ctx, w))
	})
}
