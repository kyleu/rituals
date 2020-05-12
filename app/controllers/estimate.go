package controllers

import (
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

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
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimate")
		return tmpl(templates.EstimateList(sessions, ctx, w))
	})
}

func EstimateNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := strings.TrimSpace(r.Form.Get("title"))
		if title == "" {
			title = "Untitled"
		}
		sess, err := ctx.App.Estimate.NewSession(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating estimate session"))
		}
		return ctx.Route(util.SvcEstimate, "key", sess.Slug), nil
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
			return ctx.Route("estimate.list"), nil
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("estimate.list"), "estimate")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcEstimate, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.EstimateWorkspace(sess, ctx, w))
	})
}
