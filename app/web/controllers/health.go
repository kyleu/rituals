package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web"
	"github.com/kyleu/rituals.dev/app/web/act"
)

func Health(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_, _ = w.Write([]byte("OK"))
		return "", nil
	})
}
