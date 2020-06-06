package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/kyleu/rituals.dev/app/model/auth"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/web"
)

func AuthSubmit(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func( ctx *web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		prv := auth.ProviderFromString(mux.Vars(r)[util.KeyKey])
		ref := r.Header.Get("Referer")
		secure := strings.HasSuffix(r.Proto, "s")
		state := "/"
		if len(ref) > 0 {
			u, err := url.Parse(ref)
			if err == nil && u != nil {
				state = u.Path
			}
		}

		u := ctx.App.Auth.URLFor(state, secure, prv)
		if len(u) == 0 {
			return enew(prv.Title + " is disabled")
		}
		return u, nil
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func( ctx *web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		prv := auth.ProviderFromString(mux.Vars(r)[util.KeyKey])
		code, ok := r.URL.Query()["code"]
		if !ok || len(code) == 0 {
			return enew("no auth code provided")
		}
		stateS, ok := r.URL.Query()["state"]
		u := "/"
		if ok && len(stateS) > 0 && strings.HasPrefix(stateS[0], "/") {
			u = stateS[0]
		}
		record, err := ctx.App.Auth.Handle(ctx.Profile, prv, code[0])
		if err != nil {
			return eresp(err, "")
		}

		ctx.Session.AddFlash("success:Signed in as " + record.Name)
		act.SaveSession(w, r, ctx)

		return u, nil
	})
}

func AuthSignout(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func( ctx *web.RequestContext) (string, error) {
		if !ctx.App.Auth.Enabled {
			return "", auth.ErrorAuthDisabled
		}
		id, err := act.IDFromParams(util.KeyAuth, mux.Vars(r))
		if err != nil {
			return eresp(err, util.IDErrorString(util.KeyAuth, ""))
		}

		err = ctx.App.Auth.Delete(*id)
		if err != nil {
			return eresp(err, "unable to delete auth record")
		}

		ref := r.Header.Get("Referer")
		if len(ref) == 0 {
			ref = "/"
		}

		return ref, nil
	})
}
