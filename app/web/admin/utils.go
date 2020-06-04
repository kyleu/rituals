package admin

import (
	"github.com/kyleu/rituals.dev/gen/admintemplates"
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"
)

type JSONResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Path     string    `json:"path"`
	Occurred time.Time `json:"occurred"`
}

func Modules(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyModules)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyModules, util.KeyModules)
		return tmpl(admintemplates.ModulesList(ctx, w))
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.Title(util.KeyRoutes)
		ctx.Breadcrumbs = adminBC(ctx, util.KeyRoutes, util.KeyRoutes)
		return tmpl(admintemplates.RoutesList(ctx, w))
	})
}

func adminAct(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if ctx.Profile.Role != util.RoleAdmin {
			if act.IsContentTypeJSON(act.GetContentType(r)) {
				ae := JSONResponse{Status: "error", Message: "you are not an administrator", Path: r.URL.Path, Occurred: time.Now()}
				return act.RespondJSON(w, ae, ctx.Logger)
			}
			ctx.Session.AddFlash("error:You're not an administrator, silly")
			act.SaveSession(w, r, &ctx)
			return ctx.Route("home"), nil
		}
		return f(ctx)
	})
}

func tmpl(_ int, err error) (string, error) {
	return "", err
}

func adminBC(ctx web.RequestContext, action string, name string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.AdminLink()), util.KeyAdmin)
	bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(action)), name)...)
	return bc
}

func eresp(err error, msg string) (string, error) {
	if len(msg) == 0 {
		return "", err
	}
	return "", errors.Wrap(err, msg)
}
