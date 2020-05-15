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

func RetroList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		sessions, err := ctx.App.Retro.GetByMember(ctx.Profile.UserID, 0)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving retros"))
		}

		ctx.Title = "Retrospectives"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("retro.list"), "retro")
		return tmpl(templates.RetroList(sessions, ctx, w))
	})
}

func RetroNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := strings.TrimSpace(r.Form.Get("title"))
		if title == "" {
			title = "Untitled"
		}
		sess, err := ctx.App.Retro.New(title, ctx.Profile.UserID, nil)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating retro session"))
		}
		return ctx.Route("retro", "key", sess.Slug), nil
	})
}

func RetroWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Retro.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load retro session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + key + "]")
			saveSession(w, r, ctx)
			return ctx.Route("retro.list"), nil
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("retro.list"), "retro")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcRetro, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.RetroWorkspace(sess, ctx, w))
	})
}
