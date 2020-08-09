package admin

import (
	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"net/http"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "User List"
		ctx.Breadcrumbs = adminBC(ctx, npncore.KeyUser, npncore.Plural(npncore.KeyUser))

		params := npnweb.ParamSetFromRequest(r)
		users := ctx.App.User().List(params.Get(npncore.KeyUser, ctx.Logger))
		return npncontroller.T(admintemplates.UserList(users, params, ctx, w))
	})
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		userID, err := npnweb.IDFromParams(npncore.KeyUser, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		u := ctx.App.User().GetByID(*userID, false)
		if u == nil {
			msg := "can't load user [" + userID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(npncore.KeyUser), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		auths := ctx.App.Auth().GetByUserID(*userID, params.Get(npncore.KeyAuth, ctx.Logger))
		teams := app.Team(ctx.App).GetByMember(*userID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Sprint(ctx.App).GetByMember(*userID, params.Get(util.SvcSprint.Key, ctx.Logger))
		estimates := app.Estimate(ctx.App).GetByMember(*userID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := app.Standup(ctx.App).GetByMember(*userID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := app.Retro(ctx.App).GetByMember(*userID, params.Get(util.SvcRetro.Key, ctx.Logger))
		actions := app.Action(ctx.App).GetByUser(*userID, params.Get(npncore.KeyAction, ctx.Logger))

		ctx.Title = u.Name
		bc := adminBC(ctx, npncore.KeyUser, npncore.Plural(npncore.KeyUser))
		bc = append(bc, npnweb.BreadcrumbSelf(u.Name))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.UserDetail(u, auths, teams, sprints, estimates, standups, retros, actions, params, ctx, w))
	})
}
