package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		params := paramSetFromRequest(r)
		sessions, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving teams"))
		}

		ctx.Title = "Teams"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcTeam.Key+".list"), util.SvcTeam.Key)
		return tmpl(templates.TeamList(sessions, ctx, w))
	})
}

func TeamNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		sess, err := ctx.App.Team.New(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating team session"))
		}
		return ctx.Route(util.SvcTeam.Key, "key", sess.Slug), nil
	})
}

func TeamWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sess, err := ctx.App.Team.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load team session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + key + "]")
			saveSession(w, r, ctx)
			return ctx.Route(util.SvcTeam.Key + ".list"), nil
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcTeam.Key+".list"), util.SvcTeam.Key)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcTeam.Key, "key", key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.TeamWorkspace(sess, ctx, w))
	})
}
