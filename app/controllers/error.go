package controllers

import (
	"fmt"
	"github.com/kyleu/rituals.dev/app/util"
	"net/http"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := web.ExtractContext(w, r)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, "not found")
	args := map[string]interface{}{util.KeyStatus: http.StatusInternalServerError}
	ctx.Logger.Info(fmt.Sprintf("[%v %v] returned [%d]", r.Method, r.URL.Path, http.StatusNotFound), args)
	_, _ = templates.NotFound(r, ctx, w)
}
