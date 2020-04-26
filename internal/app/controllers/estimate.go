package controllers

import (
	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func EstimateList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Estimation Sessions"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimates")
		return templates.EstimateList(ctx, w)
	})
}

func EstimateNewForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "New Estimation Session"
		bc := web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimates")
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("estimate.new.form"), "new")...)
		model := estimate.NewSession("", "", ctx.Profile.UserID)
		return templates.EstimateForm(&model, ctx, w)
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	redir(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := r.Form.Get("title")
		sess, err := ctx.App.Estimate.New(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating session"))
		}
		return ctx.Route("estimate", "key", sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		ctx.Title = "estimate [" + key + "]"
		bc := web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimates")
		ctx.Breadcrumbs = append(bc, web.BreadcrumbsSimple(ctx.Route("estimate", "key", key), key)...)
		return templates.EstimateWorkspace(ctx, w)
	})
}

