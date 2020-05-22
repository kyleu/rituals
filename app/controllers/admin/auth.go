package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AuthList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Auth List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyAuth, "auths")
		params := act.ParamSetFromRequest(r)
		users, err := ctx.App.Auth.List(params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminAuthList(users, params, ctx, w))
	})
}

func AuthDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		authID := util.GetUUIDPointer(mux.Vars(r), util.KeyID)
		if authID == nil {
			return "", errors.New("invalid auth id")
		}
		record, err := ctx.App.Auth.GetByID(*authID)
		if err != nil {
			return "", err
		}
		if record == nil {
			ctx.Session.AddFlash("error:Can't load auth [" + authID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.KeyAuth)), nil
		}

		user, err := ctx.App.User.GetByID(record.UserID, false)
		if err != nil {
			return "", err
		}
		if user == nil {
			ctx.Session.AddFlash("error:Can't load user [" + record.UserID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.KeyAuth)), nil
		}

		ctx.Title = user.Name
		bc := adminBC(ctx, util.KeyAuth, "auths")
		link := util.AdminLink(util.KeyAuth, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, authID.String()), authID.String()[0:8])...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminAuthDetail(record, user, ctx, w))
	})
}
