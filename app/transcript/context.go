package transcript

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/web"
	"logur.dev/logur"
)

type Context struct {
	UserID uuid.UUID
	App    *config.AppInfo
	Logger logur.Logger
	Routes *mux.Router
}

func (r *Context) Route(act string, pairs ...string) string {
	return r.App.Auth.FullURL(web.Route(r.Routes, r.Logger, act, pairs...))
}
