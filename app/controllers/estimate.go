package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving estimates"))
		}

		ctx.Title = util.SvcEstimate.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key+".list"), util.SvcEstimate.Key)
		return tmpl(templates.EstimateList(sessions, ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		sprintID := getUUID(r.Form, util.SvcSprint.Key)
		sess, err := ctx.App.Estimate.New(title, ctx.Profile.UserID, teamID, sprintID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating estimate session"))
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot send content update"))
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, sprintID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot send content update"))
		}
		return ctx.Route(util.SvcEstimate.Key, util.KeyKey, sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Estimate.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load estimate session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcEstimate.Key + ".list"), nil
		}

		var tm *team.Session
		var tmTitle string
		if sess.TeamID != nil {
			tm, _ = ctx.App.Team.GetByID(*sess.TeamID)
			tmTitle = tm.Title
		}

		var spr *sprint.Session
		if sess.SprintID != nil {
			spr, _ = ctx.App.Sprint.GetByID(*sess.SprintID)
		}

		auths, currTeams, err := authsAndTeams(ctx, sess.TeamID)
		permErrors := ctx.App.Estimate.Permissions.Check(util.SvcEstimate, sess.ID, auths, sess.TeamID, tmTitle, currTeams)
		if len(permErrors) > 0 {
			return permErrorTemplate(permErrors, ctx, w)
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key+".list"), util.SvcEstimate.Key)
		if spr != nil {
			bc = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, spr.Slug), spr.Title)
		}
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.EstimateWorkspace(sess, ctx, w))
	})
}
