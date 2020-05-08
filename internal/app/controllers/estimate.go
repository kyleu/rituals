package controllers

import (
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/internal/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/internal/app/estimate"
	"github.com/kyleu/rituals.dev/internal/app/web"

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
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("estimate.new.form"), "new")...)
		ctx.Breadcrumbs = bc

		model := estimate.NewSession("", "", ctx.Profile.UserID)
		return templates.EstimateForm(&model, ctx, w)
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	redir(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := strings.TrimSpace(r.Form.Get("title"))
		if title == "" {
			title = "Untitled"
		}
		sess, err := ctx.App.Estimate.NewSession(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating session"))
		}
		return ctx.Route(util.SvcEstimate, "key", sess.Slug), nil
	})
}

func EstimateWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		est, err := ctx.App.Estimate.GetBySlug(key)
		if err != nil {
			return 0, errors.WithStack(errors.Wrap(err, "cannot load session"))
		}

		ctx.Title = est.Title
		bc := web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimates")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate, "key", key), est.Title)...)
		ctx.Breadcrumbs = bc

		return templates.EstimateWorkspace(est, ctx, w)
	})
}
