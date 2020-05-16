package controllers

import (
	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		sessions, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, 0)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving sprints"))
		}

		ctx.Title = "Daily Sprints"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key + ".list"), "sprint")
		return tmpl(templates.SprintList(sessions, ctx, w))
	})
}

func SprintNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		sess, err := ctx.App.Sprint.New(title, ctx.Profile.UserID, nil)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating sprint session"))
		}
		return ctx.Route(util.SvcSprint.Key, "key", sess.Slug), nil
	})
}

func SprintWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Sprint.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load sprint session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + key + "]")
			saveSession(w, r, ctx)
			return ctx.Route(util.SvcSprint.Key + ".list"), nil
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key + ".list"), "sprint")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.SprintWorkspace(sess, ctx, w))
	})
}
