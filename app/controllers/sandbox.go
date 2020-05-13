package controllers

import (
	"net/http"

	web "github.com/kyleu/rituals.dev/app/web"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/gen/templates"
)

var _sandboxes = []string{"gallery", "parse", "testbed"}

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sandboxes"
		ctx.Breadcrumbs = append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		return tmpl(templates.SandboxList(_sandboxes, ctx, w))
	})
}

func SandboxForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		if key == "testbed" {
			return "", errors.WithStack(errors.New("error!"))
		}
		ctx.Title = "[" + key + "] Sandbox"
		bc := append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		bc = append(bc, web.Breadcrumb{Path: ctx.Route("sandbox.run", "key", key), Title: key})
		ctx.Breadcrumbs = bc

		return tmpl(templates.SandboxForm(key, ctx, w))
	})
}
