package transcript

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/kyleu/npn/npnweb"
	"logur.dev/logur"
)

type Context struct {
	UserID uuid.UUID
	App    npnweb.AppInfo
	Logger logur.Logger
	Routes *mux.Router
}

func (r *Context) Route(act string, pairs ...string) string {
	return r.App.Auth().FullURL(npnweb.Route(r.Routes, r.Logger, act, pairs...))
}
