package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"net/url"
	"time"

	"emperror.dev/errors"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
	"golang.org/x/text/language"

	"github.com/kyleu/rituals.dev/app/util"
)

func act(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	startNanos := time.Now().UnixNano()
	ctx := web.ExtractContext(w, r)

	if len(ctx.Flashes) > 0 {
		saveSession(w, r, ctx)
	}

	redir, err := f(ctx)
	if err != nil {
		ctx.Logger.Warn(fmt.Sprintf("error running action: %+v", errors.WithStack(err)))
		if ctx.Title == "" {
			ctx.Title = "Error"
		}
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/json", "text/json":
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			b, _ := json.MarshalIndent(ErrorResult{Status: "error", Message: err.Error()}, "", "  ")
			_, _ = w.Write(b)
		default:
			_, _ = templates.InternalServerError(util.GetErrorDetail(err), r, ctx, w)
		}
	}
	if redir != "" {
		if len(ctx.Flashes) > 0 {
			saveSession(w, r, ctx)
		}
		w.Header().Set("Location", redir)
		w.WriteHeader(http.StatusFound)
		logComplete(startNanos, ctx, http.StatusFound, r)
	} else {
		logComplete(startNanos, ctx, http.StatusOK, r)
	}
}

func adminAct(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		if ctx.Profile.Role != util.RoleAdmin {
			ctx.Session.AddFlash("error:You're not an administrator, silly")
			saveSession(w, r, ctx)
			return ctx.Route("home"), nil
		}
		return f(ctx)
	})
}

func tmpl(_ int, err error) (string, error) {
	return "", err
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
		ctx.Logger.Warn("unable to save session to response")
	}
}

func getSprintID(form url.Values) *uuid.UUID {
	sprintString := form.Get(util.SvcSprint.Key)
	var sprintID *uuid.UUID
	if sprintString != "" {
		s, err := uuid.FromString(sprintString)
		if err == nil {
			sprintID = &s
		}
	}
	return sprintID
}
