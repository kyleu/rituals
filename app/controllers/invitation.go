package controllers

import (
	"net/http"

	web "github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func JoinPost(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Join Session"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.post"), "join")
		return tmpl(templates.Todo("JoinPost", ctx, w))
	})
}

func JoinGet(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		ctx.Title = "[" + key + "]"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.form", "key", key), "join")
		return tmpl(templates.Todo("JoinGet", ctx, w))
	})
}
