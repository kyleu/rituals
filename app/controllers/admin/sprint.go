package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sprint List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)

		params := act.ParamSetFromRequest(r)
		sprints := ctx.App.Sprint.List(params.Get(util.SvcSprint.Key, ctx.Logger))
		return tmpl(templates.AdminSprintList(sprints, params, ctx, w))
	})
}

func SprintDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		sprintID, err := act.IDFromParams(util.SvcSprint.Key, mux.Vars(r))
		if err != nil {
			return eresp(err, "")
		}
		sess, err := ctx.App.Sprint.GetByID(*sprintID)
		if err != nil {
			return eresp(err, "")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + sprintID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcSprint.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		estimates := ctx.App.Estimate.GetBySprint(*sprintID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		standups := ctx.App.Standup.GetBySprint(*sprintID, params.Get(util.SvcStandup.Key, ctx.Logger))
		retros := ctx.App.Retro.GetBySprint(*sprintID, params.Get(util.SvcRetro.Key, ctx.Logger))

		comments := ctx.App.Sprint.Comments.GetByModelID(*sprintID, params.Get(util.KeyComment, ctx.Logger))
		actions := ctx.App.Action.GetBySvcModel(util.SvcSprint, *sprintID, params.Get(util.KeyAction, ctx.Logger))
		members := ctx.App.Sprint.Members.GetByModelID(*sprintID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Sprint.Permissions.GetByModelID(*sprintID, params.Get(util.KeyPermission, ctx.Logger))

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)
		link := util.AdminLink(util.SvcSprint.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, sprintID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminSprintDetail(sess, estimates, standups, retros, comments, members, perms, actions, params, ctx, w))
	})
}
