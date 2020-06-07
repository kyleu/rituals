package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		ctx.Title = util.SvcStandup.Title + " List"

		ctx.Breadcrumbs = adminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)

		params := act.ParamSetFromRequest(r)
		standups := ctx.App.Standup.List(params.Get(util.SvcStandup.Key, ctx.Logger))
		return tmpl(admintemplates.StandupList(standups, params, ctx, w))
	})
}

func StandupDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *web.RequestContext) (string, error) {
		standupID, err := act.IDFromParams(util.SvcStandup.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess := ctx.App.Standup.GetByID(*standupID)
		if sess == nil {
			msg := "can't load standup [" + standupID.String() + "]"
			return act.FlashAndRedir(false, msg, util.AdminLink(util.SvcStandup.Key), w, r, ctx)
		}

		params := act.ParamSetFromRequest(r)

		reports := ctx.App.Standup.GetReports(*standupID, params.Get(util.KeyReport, ctx.Logger))

		data := ctx.App.Standup.Data.GetData(*standupID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)
		link := util.AdminLink(util.SvcStandup.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, standupID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.StandupDetail(sess, reports, data, params, ctx, w))
	})
}
