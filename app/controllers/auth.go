package controllers

import (
	"github.com/kyleu/rituals.dev/app/auth"
	"net/http"
	"net/url"
	"strings"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"emperror.dev/errors"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/web"
)

func AuthSubmit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		key := mux.Vars(r)[util.KeyKey]
		host := r.Header.Get("Host")
		ref := r.Header.Get("Referer")
		secure := strings.HasSuffix(r.Proto, "s")
		state := "/"
		if len(ref) > 0 {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				state = u.Path
			}
		}

		u := ctx.App.Auth.URLFor(state, secure, host, key)
		if len(u) == 0 {
			return "", util.IDError(util.KeyAuth, key)
		}
		return u, nil
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		key := mux.Vars(r)[util.KeyKey]
		code, ok := r.URL.Query()["code"]
		if !ok || len(code) == 0 {
			return "", errors.New("no auth code provided")
		}
		stateS, ok := r.URL.Query()["state"]
		u := "/"
		if ok && len(stateS) > 0 && strings.HasPrefix(stateS[0], "/") {
			u = stateS[0]
		}
		record, err := ctx.App.Auth.Handle(ctx.Profile, key, code[0])
		if err != nil {
			return "", err
		}

		ctx.Session.AddFlash("success:Signed in as " + record.Name)
		act.SaveSession(w, r, ctx)


		return u, nil
	})
}
