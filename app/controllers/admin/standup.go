package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcStandup.Title + " List"

		ctx.Breadcrumbs = adminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)

		params := act.ParamSetFromRequest(r)
		standups := ctx.App.Standup.List(params.Get(util.SvcStandup.Key, ctx.Logger))
		return tmpl(templates.AdminStandupList(standups, params, ctx, w))
	})
}

func StandupDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		standupID, err := act.IDFromParams(util.SvcStandup.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess, err := ctx.App.Standup.GetByID(*standupID)
		if err != nil {
			return eresp(err, "")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + standupID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcStandup.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		reports := ctx.App.Standup.GetReports(*standupID, params.Get(util.KeyReport, ctx.Logger))
		comments := ctx.App.Standup.Comments.GetByModelID(*standupID, params.Get(util.KeyComment, ctx.Logger))
		actions := ctx.App.Action.GetBySvcModel(util.SvcStandup, *standupID, params.Get(util.KeyAction, ctx.Logger))
		members := ctx.App.Standup.Members.GetByModelID(*standupID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Standup.Permissions.GetByModelID(*standupID, params.Get(util.KeyPermission, ctx.Logger))


		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcStandup.Key, util.SvcStandup.Plural)
		link := util.AdminLink(util.SvcStandup.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, standupID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminStandupDetail(sess, reports, comments, members, perms, actions, params, ctx, w))
	})
}
