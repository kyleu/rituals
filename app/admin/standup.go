package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = util.SvcStandup.Title + " List"

		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)

		params := npnweb.ParamSetFromRequest(r)
		standups := app.Standup(ctx.App).List(params.Get(util.SvcStandup.Key, ctx.Logger))
		return npncontroller.T(admintemplates.StandupList(standups, params, ctx, w))
	})
}

func StandupDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		standupID, err := npnweb.IDFromParams(util.SvcStandup.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Standup(ctx.App).GetByID(*standupID)
		if sess == nil {
			msg := "can't load standup [" + standupID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcStandup.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		reports := app.Standup(ctx.App).GetReports(*standupID, params.Get(npncore.KeyReport, ctx.Logger))

		data := app.Standup(ctx.App).Data.GetData(*standupID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := npncontroller.AdminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.StandupDetail(sess, reports, data, params, ctx, w))
	})
}
