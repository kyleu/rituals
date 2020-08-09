package admin

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"net/http"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/email"
	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/gorilla/mux"

)

func EmailList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Email List"
		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyEmail, npncore.Plural(npncore.KeyEmail))

		params := npnweb.ParamSetFromRequest(r)
		emailSvc := email.NewService(app.Database(ctx.App), ctx.Logger)
		emails := emailSvc.List(params.Get(npncore.KeyEmail, ctx.Logger))
		return npncontroller.T(admintemplates.EmailList(emails, params, ctx, w))
	})
}

func EmailDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		emailID, ok := mux.Vars(r)[npncore.KeyID]
		if !ok {
			return npncontroller.EResp(errors.New("invalid email id"))
		}
		emailSvc := email.NewService(app.Database(ctx.App), ctx.Logger)
		e := emailSvc.GetByID(emailID)
		if e == nil {
			msg := "can't load email [" + emailID + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyEmail), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		ctx.Title = e.ID
		bc := adminBC(ctx, npncore.KeyEmail, npncore.Plural(npncore.KeyEmail))
		bc = append(bc, npnweb.BreadcrumbSelf(e.ID))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.EmailDetail(e, params, ctx, w))
	})
}
