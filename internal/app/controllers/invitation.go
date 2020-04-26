package controllers

import (
	"net/http"

	web "github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/internal/gen/templates"
)

func JoinForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Join Session"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.form"), "join")
		return templates.InvitationForm("", ctx, w)
	})
}

func JoinPost(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Join Session"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.post"), "join")
		return templates.Todo("JoinPost", ctx, w)
	})
}

func JoinGet(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		ctx.Title = "[" + key + "]"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.form", "key", key), "join")
		return templates.Todo("JoinGet", ctx, w)
	})
}

func SessionNew(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "New Session"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("session.new"), "new")
		return templates.Todo("SessionNew", ctx, w)
	})
}
