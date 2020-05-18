package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/web"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/gen/templates"
)

type ErrorResult struct {
	Status  string
	Message string
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	ctx := web.ExtractContext(w, r)
	ctx.Title = "Not Found"
	ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, "not found")
	args := map[string]interface{}{"status": 500}
	ctx.Logger.Info(fmt.Sprintf("[%v %v] returned [%d]", r.Method, r.URL.Path, http.StatusNotFound), args)
	_, _ = templates.NotFound(r, ctx, w)
}

func internalServerError(router *mux.Router, info *config.AppInfo, w http.ResponseWriter, r *http.Request) {
	defer lastChanceError(w)

	if err := recover(); err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		rc := context.WithValue(r.Context(), routesKey, router)
		rc = context.WithValue(rc, infoKey, info)
		ctx := web.ExtractContext(w, r.WithContext(rc))
		ctx.Title = "Server Error"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(r.URL.Path, "error")
		e, ok := err.(error)
		if !ok {
			e = errors.New(fmt.Sprintf("err [%v] is of type [%T]", err, err))
		}
		_, _ = templates.InternalServerError(util.GetErrorDetail(e), r, ctx, w)
		args := map[string]interface{}{"status": 500}
		st := http.StatusInternalServerError
		ctx.Logger.Warn(fmt.Sprintf("[%v %v] returned [%d]: %+v", r.Method, r.URL.Path, st, e), args)
	}
}

func lastChanceError(w io.Writer) {
	if err := recover(); err != nil {
		println(fmt.Sprintf("error while processing error handler: %+v", err))
		_, _ = w.Write([]byte("Internal Server Error"))
	}
}
