package act

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/gen/templates"
	"golang.org/x/text/language"
)

type errorResult struct {
	Status  string
	Message string
}

func Act(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	startNanos := time.Now().UnixNano()
	ctx := web.ExtractContext(w, r, false)

	if len(ctx.Flashes) > 0 {
		SaveSession(w, r, ctx)
	}

	redir, err := f(ctx)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("%+v", errors.Wrap(err, "error running action")))
		if ctx.Title == "" {
			ctx.Title = "Error"
		}
		contentType := r.Header.Get("Content-Type")

		switch contentType {
		case "application/json", "text/json":
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			b, _ := json.MarshalIndent(errorResult{Status: "error", Message: err.Error()}, "", "  ")
			_, _ = w.Write(b)
		default:
			_, _ = templates.InternalServerError(util.GetErrorDetail(err), r, ctx, w)
		}
	}
	if redir != "" {
		if len(ctx.Flashes) > 0 {
			SaveSession(w, r, ctx)
		}
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
		logComplete(startNanos, ctx, http.StatusFound, r)
	} else {
		logComplete(startNanos, ctx, http.StatusOK, r)
	}
}

func SaveSession(w http.ResponseWriter, r *http.Request, ctx web.RequestContext) {
	err := ctx.Session.Save(r, w)
	if err != nil {
		ctx.Logger.Warn("unable to save session to response")
	}
}

func logComplete(startNanos int64, ctx web.RequestContext, status int, r *http.Request) {
	delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := util.MicrosToMillis(language.AmericanEnglish, int(delta))
	args := map[string]interface{}{"elapsed": delta, util.KeyStatus: status}
	ctx.Logger.Debug(fmt.Sprintf("[%v %v] returned [%v] in [%v]", r.Method, r.URL.Path, status, ms), args)
}
