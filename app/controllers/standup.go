package controllers

import (
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		sessions, err := ctx.App.Standup.GetByMember(ctx.Profile.UserID, 0)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving standups"))
		}

		ctx.Title = "Daily Standups"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("standup.list"), "standup")
		return tmpl(templates.StandupList(sessions, ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := strings.TrimSpace(r.Form.Get("title"))
		if title == "" {
			title = "Untitled"
		}
		sess, err := ctx.App.Standup.NewSession(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating standup session"))
		}
		return ctx.Route("standup", "key", sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Standup.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load standup session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + key + "]")
			saveSession(w, r, ctx)
			return ctx.Route("standup.list"), nil
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("standup.list"), "standup")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcStandup, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.StandupWorkspace(sess, ctx, w))
	})
}
