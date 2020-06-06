package admin

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/model/email"
	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func EmailList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = "Email List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyEmail, util.Plural(util.KeyEmail))

		params := act.ParamSetFromRequest(r)
		emailSvc := email.NewService(ctx.App.Database, ctx.Logger)
		emails := emailSvc.List(params.Get(util.KeyEmail, ctx.Logger))
		return tmpl(admintemplates.EmailList(emails, params, ctx, w))
	})
}

func EmailDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		emailID, ok := mux.Vars(r)[util.KeyID]
		if !ok {
			return eresp(errors.New("invalid email id"), "")
		}
		emailSvc := email.NewService(ctx.App.Database, ctx.Logger)
		e := emailSvc.GetByID(emailID)
		if e == nil {
			ctx.Session.AddFlash("error:Can't load email [" + emailID + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.KeyEmail)), nil
		}

		params := act.ParamSetFromRequest(r)

		ctx.Title = e.ID
		bc := adminBC(ctx, util.KeyEmail, util.Plural(util.KeyEmail))
		link := util.AdminLink(util.KeyEmail, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, e.ID), e.ID)...)
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.EmailDetail(e, params, ctx, w))
	})
}
