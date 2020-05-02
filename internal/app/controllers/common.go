package controllers

import (
	"fmt"
	"net/http"
	"time"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/kyleu/rituals.dev/internal/gen/templates"
	"golang.org/x/text/language"

	"github.com/kyleu/rituals.dev/internal/app/util"
)

func act(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (int, error)) {
	startNanos := time.Now().UnixNano()
	ctx := web.ExtractContext(w, r)

	if len(ctx.Flashes) > 0 {
		saveSession(w, r, ctx)
	}

	_, err := f(ctx)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("error running action: %+v", errors.WithStack(err)))
		if ctx.Title == "" {
			ctx.Title = "Error"
		}
		_, _ = templates.InternalServerError(util.GetErrorDetail(err), r, ctx, w)
	}
	logComplete(startNanos, ctx, http.StatusOK, r)
}

func redir(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	startNanos := time.Now().UnixNano()
	ctx := web.ExtractContext(w, r)
	url, err := f(ctx)
	if err == nil {
		if len(ctx.Flashes) > 0 {
			saveSession(w, r, ctx)
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		logComplete(startNanos, ctx, http.StatusFound, r)
	} else {
		ctx.Logger.Warn(fmt.Sprintf("error running redirect: %+v", errors.WithStack(err)))
		if ctx.Title == "" {
			ctx.Title = "Error"
		}
		_, _ = templates.InternalServerError(util.GetErrorDetail(err), r, ctx, w)
	}
}

func logComplete(startNanos int64, ctx web.RequestContext, status int, r *http.Request) {
	delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := util.MicrosToMillis(language.AmericanEnglish, int(delta))
	args := map[string]interface{}{"elapsed": delta, "status": status}
	ctx.Logger.Debug(fmt.Sprintf("[%v %v] returned [%v] in [%v]", r.Method, r.URL.Path, status, ms), args)
}

func saveSession(w http.ResponseWriter, r *http.Request, ctx web.RequestContext) {
	err := ctx.Session.Save(r, w)
	if err != nil {
		ctx.Logger.Warn("Unable to save session to response")
	}
}
