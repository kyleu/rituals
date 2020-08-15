package admin

import (
	"net/http"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/kyleu/rituals.dev/gen/admintemplates"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		ctx.Title = "Sprint List"
		ctx.Breadcrumbs = npncontroller.AdminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)

		params := npnweb.ParamSetFromRequest(r)
		sprints := app.Sprint(ctx.App).List(params.Get(util.SvcSprint.Key, ctx.Logger))
		return npncontroller.T(admintemplates.SprintList(sprints, params, ctx, w))
	})
}

func SprintDetail(w http.ResponseWriter, r *http.Request) {
	npncontroller.AdminAct(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		sprintID, err := npnweb.IDFromParams(util.SvcSprint.Key, mux.Vars(r))
		if err != nil {
			return npncontroller.EResp(err)
		}
		sess := app.Sprint(ctx.App).GetByID(*sprintID)
		if sess == nil {
			msg := "can't load sprint [" + sprintID.String() + "]"
			return npncontroller.FlashAndRedir(false, msg, npnweb.AdminLink(util.SvcSprint.Key), w, r, ctx)
		}

		params := npnweb.ParamSetFromRequest(r)

		estimates := app.Estimate(ctx.App).GetBySprintID(*sprintID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := app.Standup(ctx.App).GetBySprintID(*sprintID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := app.Retro(ctx.App).GetBySprintID(*sprintID, params.Get(util.SvcRetro.Key, ctx.Logger))

		data := app.Sprint(ctx.App).Data.GetData(*sprintID, params, ctx.Logger)

		ctx.Title = sess.Title
		bc := npncontroller.AdminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)
		bc = append(bc, npnweb.BreadcrumbSelf(sess.Slug))
		ctx.Breadcrumbs = bc

		return npncontroller.T(admintemplates.SprintDetail(sess, estimates, standups, retros, data, params, ctx, w))
	})
}
