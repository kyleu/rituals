package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Retrospectives"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("retro.list"), "retros")
		return templates.RetroList(ctx, w)
	})
}

func RetroNewForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "New Retrospective"
		bc := web.BreadcrumbsSimple(ctx.Route("retro.list"), "retros")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("retro.new.form"), "new")...)
		ctx.Breadcrumbs = bc

		return templates.Todo("New retrospective!", ctx, w)
	})
}

func RetroNew(w http.ResponseWriter, r *http.Request) {
	redir(w, r, func(ctx web.RequestContext) (string, error) {
		newKey := "todo"
		return ctx.Route("retro", "key", newKey), nil
	})
}

func RetroWorkspace(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		ctx.Title = "retro [" + key + "]"
		bc := web.BreadcrumbsSimple(ctx.Route("retro.list"), "retros")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("retro", "key", key), key)...)
		ctx.Breadcrumbs = bc

		return templates.RetroWorkspace(ctx, w)
	})
}
