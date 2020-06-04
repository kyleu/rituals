package admin

import (
	"github.com/kyleu/rituals.dev/gen/admintemplates"
	"net/http"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

)

func UserList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "User List"
		ctx.Breadcrumbs = adminBC(ctx, util.KeyUser, util.Plural(util.KeyUser))

		params := act.ParamSetFromRequest(r)
		users := ctx.App.User.List(params.Get(util.KeyUser, ctx.Logger))
		return tmpl(admintemplates.UserList(users, params, ctx, w))
	})
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		userID, err := act.IDFromParams(util.KeyUser, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		u := ctx.App.User.GetByID(*userID, false)
		if u == nil {
			ctx.Session.AddFlash("error:Can't load user [" + userID.String() + "]")
			act.SaveSession(w, r, &ctx)
			return ctx.Route(util.AdminLink(util.KeyUser)), nil
		}

		params := act.ParamSetFromRequest(r)

		auths := ctx.App.Auth.GetByUserID(*userID, params.Get(util.KeyAuth, ctx.Logger))
		teams := ctx.App.Team.GetByMember(*userID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := ctx.App.Sprint.GetByMember(*userID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := ctx.App.Estimate.GetByMember(*userID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := ctx.App.Standup.GetByMember(*userID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := ctx.App.Retro.GetByMember(*userID, params.Get(util.SvcRetro.Key, ctx.Logger))
		actions := ctx.App.Action.GetByUser(*userID, params.Get(util.KeyAction, ctx.Logger))

		ctx.Title = u.Name
		bc := adminBC(ctx, util.KeyUser, util.Plural(util.KeyUser))
		link := util.AdminLink(util.KeyUser, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, userID.String()), u.Name)...)
		ctx.Breadcrumbs = bc

		return tmpl(admintemplates.UserDetail(u, auths, teams, sprints, estimates, standups, retros, actions, params, ctx, w))
	})
}
