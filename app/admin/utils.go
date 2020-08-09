package admin

import (
	"net/http"
	"time"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnweb"
)

type JSONResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Path     string    `json:"path"`
	Occurred time.Time `json:"occurred"`
}

func adminAct(w http.ResponseWriter, r *http.Request, f func(*npnweb.RequestContext) (string, error)) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		if ctx.Profile.Role != npnuser.RoleAdmin {
			if npncontroller.IsContentTypeJSON(npncontroller.GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: "you are not an administrator", Path: r.URL.Path, Occurred: time.Now()}
				return npncontroller.RespondJSON(w, "", ae, ctx.Logger)
			}
			msg := "you're not an administrator, silly!"
			return npncontroller.FlashAndRedir(false, msg, "home", w, r, ctx)
		}
		return f(ctx)
	})
}

func adminBC(ctx *npnweb.RequestContext, action string, name string) npnweb.Breadcrumbs {
	bc := npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink()), npncore.KeyAdmin)
	bc = append(bc, npnweb.BreadcrumbsSimple(ctx.Route(npnweb.AdminLink(action)), name)...)
	return bc
}
