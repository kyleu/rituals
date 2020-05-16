package controllers

import (
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		sessions, err := ctx.App.Estimate.GetByMember(ctx.Profile.UserID, 0)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving estimates"))
		}

		ctx.Title = "Estimation Sessions"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key + ".list"), util.SvcEstimate.Key)
		return tmpl(templates.EstimateList(sessions, ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		sess, err := ctx.App.Estimate.New(title, ctx.Profile.UserID, nil)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating estimate session"))
		}
		return ctx.Route(util.SvcEstimate.Key, "key", sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Estimate.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load estimate session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load estimate [" + key + "]")
			saveSession(w, r, ctx)
			return ctx.Route(util.SvcEstimate.Key + ".list"), nil
		}

		var spr *sprint.Session
		if sess.SprintID != nil {
			spr, _ = ctx.App.Sprint.GetByID(*sess.SprintID)
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key + ".list"), util.SvcEstimate.Key)
		if spr != nil {
			bc = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, "key", spr.Slug), spr.Title)
		}
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate.Key, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.EstimateWorkspace(sess, spr, ctx, w))
	})
}
