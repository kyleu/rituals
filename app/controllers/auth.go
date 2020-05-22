package controllers

import (
	"net/http"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/web"
)

func AuthSubmit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		host := r.Header.Get("Host")
		secure := strings.HasSuffix(r.Proto, "s")
		url := ctx.App.Auth.URLFor(secure, host, key)
		if len(url) == 0 {
			return "", errors.New("invalid auth key [" + key + "]")
		}
		return url, nil
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		code, ok := r.URL.Query()["code"]
		if !ok || len(code) == 0 {
			return "", errors.New("no auth code provided")
		}
		record, err := ctx.App.Auth.Handle(ctx.Profile, key, code[0])
		if err != nil {
			return "", err
		}

		ctx.Session.AddFlash("success:Signed in as " + record.Name)
		act.SaveSession(w, r, ctx)

		return ctx.Route("home"), nil
	})
}
