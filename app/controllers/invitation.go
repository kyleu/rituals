package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func JoinPost(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Join Session"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.post"), "join")
		return tmpl(templates.Message("JoinPost", ctx, w))
	})
}

func JoinGet(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		ctx.Title = "[" + key + "]"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("join.form", util.KeyKey, key), "join")
		return tmpl(templates.Message("JoinGet", ctx, w))
	})
}
