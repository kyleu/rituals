package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Daily Standups"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("standup.list"), "standups")
		return templates.StandupList(ctx, w)
	})
}

func StandupNewForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "New Daily Standup"
		bc := web.BreadcrumbsSimple(ctx.Route("standup.list"), "standups")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("standup.new.form"), "new")...)
		ctx.Breadcrumbs = bc

		return templates.Todo("New daily standup!", ctx, w)
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	redir(w, r, func(ctx web.RequestContext) (string, error) {
		newKey := "todo"
		return ctx.Route("standup", "key", newKey), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		ctx.Title = "standup [" + key + "]"
		bc := web.BreadcrumbsSimple(ctx.Route("standup.list"), "standups")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("standup", "key", key), key)...)
		ctx.Breadcrumbs = bc

		return templates.StandupWorkspace(ctx, w)
	})
}
