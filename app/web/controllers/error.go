package controllers

import (
	"fmt"
	"github.com/kyleu/rituals.dev/gen/components"
	"net/http"

	"github.com/kyleu/rituals.dev/app/web"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := web.ExtractContext(w, r, false)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, "not found")
	ctx.Logger.Info(fmt.Sprintf("[%v %v] returned [%d]", r.Method, r.URL.Path, http.StatusNotFound))
	_, _ = components.NotFound(r, ctx, w)
}
