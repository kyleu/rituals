package act

import (
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/gorilla/sessions"
	"logur.dev/logur"

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
		util.LogWarn(ctx.Logger, "error running action: %+v", err)
		if ctx.Title == "" {
			ctx.Title = "Error"
		}
		if IsContentTypeJSON(GetContentType(r)) {
			_, _ = RespondJSON(w, errorResult{Status: util.KeyError, Message: err.Error()}, ctx.Logger)
		} else {
			_, _ = templates.InternalServerError(util.GetErrorDetail(err), r, ctx, w)
		}
	}
	if redir != "" {
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
		logComplete(startNanos, ctx, http.StatusFound, r)
	} else {
		logComplete(startNanos, ctx, http.StatusOK, r)
	}
}

func RespondJSON(w http.ResponseWriter, body interface{}, logger logur.Logger) (string, error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	b := util.ToJSONBytes(body, logger)
	_, err := w.Write(b)
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	return "", nil
}

func SaveSession(w http.ResponseWriter, r *http.Request, ctx web.RequestContext) {
	ctx.Session.Options = &sessions.Options{Path: "/", HttpOnly: true, SameSite: http.SameSiteDefaultMode}
	err := ctx.Session.Save(r, w)
	if err != nil {
		ctx.Logger.Warn("unable to save session to response")
	}
}

func logComplete(startNanos int64, ctx web.RequestContext, status int, r *http.Request) {
	delta := (time.Now().UnixNano() - startNanos) / int64(time.Microsecond)
	ms := util.MicrosToMillis(language.AmericanEnglish, int(delta))
	// args := map[string]interface{}{"elapsed": delta, util.KeyStatus: status}
	util.LogDebug(ctx.Logger, "[%v %v] returned [%v] in [%v]", r.Method, r.URL.Path, status, ms)
}
