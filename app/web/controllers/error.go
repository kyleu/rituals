package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := web.ExtractContext(w, r, false)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, "not found")
	util.LogInfo(ctx.Logger, "[%v %v] returned [%d]", r.Method, r.URL.Path, http.StatusNotFound)
	_, _ = templates.NotFound(r, ctx, w)
}
