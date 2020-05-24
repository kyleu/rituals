package admin

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func Home(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Admin"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.AdminLink()), util.KeyAdmin)
		return tmpl(templates.AdminHome(ctx, w))
	})
}

func adminAct(w http.ResponseWriter, r *http.Request, f func(web.RequestContext) (string, error)) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		if ctx.Profile.Role != util.RoleAdmin {
			ctx.Session.AddFlash("error:You're not an administrator, silly")
			act.SaveSession(w, r, ctx)
			return ctx.Route("home"), nil
		}
		return f(ctx)
	})
}

func tmpl(_ int, err error) (string, error) {
	return "", err
}

func adminBC(ctx web.RequestContext, action string, s string) web.Breadcrumbs {
	bc := web.BreadcrumbsSimple(ctx.Route(util.AdminLink()), util.KeyAdmin)
	bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.AdminLink(action)), s)...)
	return bc
}

func idFromParams(svc string, m map[string]string) (*uuid.UUID, error) {
	retOut, ok := m[util.KeyID]

	if !ok {
		return nil, errors.New("params do not contain \"id\"")
	}

	ret := util.GetUUIDFromString(retOut)

	if ret == nil {
		return nil, util.IDError(svc, retOut)
	}

  return ret, nil
}
