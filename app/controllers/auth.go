package controllers

import (
	"emperror.dev/errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/web"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AuthSubmit(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		url := ctx.App.Auth.UrlFor(key)
		if len(url) == 0 {
			return "", errors.New("invalid auth key [" + key + "]")
		}
		return url, nil
	})
}

func AuthCallback(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)["key"]
		code, ok := r.URL.Query()["code"]
		if !ok || len(code) == 0 {
			return "", errors.New("no auth code provided")
		}
		record, err := ctx.App.Auth.Handle(ctx.Profile, key, code[0])
		if err != nil {
			return "", err
		}

		return tmpl(templates.Todo(fmt.Sprintf("Auth Callback: %v", record), ctx, w))
	})
}
