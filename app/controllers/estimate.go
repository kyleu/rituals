package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

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

		permErrors, bc := check(&ctx, ctx.App.Estimate.Permissions, util.SvcEstimate, sess.ID, key, sess.Title, sess.TeamID, sess.SprintID)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcEstimate, permErrors, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.EstimateWorkspace(sess, ctx, w))
	})
}
