package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kyleu/rituals.dev/app/sandbox"

	web "github.com/kyleu/rituals.dev/app/web"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/gen/templates"
)

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sandboxes"
		ctx.Breadcrumbs = append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		return tmpl(templates.SandboxList(sandbox.AllSandboxes, ctx, w))
	})
}

func SandboxRun(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		sb := sandbox.SandboxFromString(key)
		if sb == nil {
			return "", errors.New("invalid sandbox [" + key + "]")
		}
		content, err := sb.Resolve(ctx)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error running sandbox ["+key+"]"))
		}

		js, err := json.Marshal(content)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error marshalling sandbox ["+key+"] response"))
		}

		ctx.Title = sb.Title + " Sandbox"
		bc := append(aboutBC(ctx), web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")...)
		bc = append(bc, web.Breadcrumb{Path: ctx.Route("sandbox.run", "key", key), Title: key})
		ctx.Breadcrumbs = bc

		return tmpl(templates.SandboxRun(sb, string(js), ctx, w))
	})
}
