package admin

import (
	"github.com/gofrs/uuid"
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/gorilla/mux"
)

func AuthList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Auth List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, npncore.KeyAuth, npncore.Plural(npncore.KeyAuth))
		params := npnweb.ParamSetFromRequest(r)
		users := ctx.App.Auth().List(params.Get(npncore.KeyAuth, ctx.Logger))
		return npncontroller.T(admintemplates.AuthList(users, params, ctx, w))
	})
}

func AuthDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		authID, err := npnweb.IDFromParams(npncore.KeyAuth, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		userID := uuid.UUID{} // TODO
		record := ctx.App.Auth().GetByID(userID, *authID)
		if record == nil {
			msg := "can't load auth [" + authID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyAuth), w, r, ctx)
		}

		user := ctx.App.User().GetByID(record.UserID, false)
		if user == nil {
			msg := "can't load user [" + record.UserID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyAuth), w, r, ctx)
		}

		ctx.Title = user.Name
		bc := npncontroller.AdminBC(ctx, npncore.KeyAuth, npncore.Plural(npncore.KeyAuth))
		bc = append(bc, npnweb.BreadcrumbSelf(authID.String()[0:8]))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.AuthDetail(record, user, ctx, w))
	})
}
