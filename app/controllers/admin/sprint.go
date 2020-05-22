package admin

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sprint List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)

		params := act.ParamSetFromRequest(r)
		sprints, err := ctx.App.Sprint.List(params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminSprintList(sprints, params, ctx, w))
	})
}

func SprintDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		sprintID := util.GetUUIDPointer(mux.Vars(r), util.KeyID)
		if sprintID == nil {
			return "", errors.New("invalid sprint id")
		}
		sess, err := ctx.App.Sprint.GetByID(*sprintID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + sprintID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcSprint.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		members := ctx.App.Sprint.Members.GetByModelID(*sprintID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Sprint.Permissions.GetByModelID(*sprintID, params.Get(util.KeyPermission, ctx.Logger))

		estimates, err := ctx.App.Estimate.GetBySprint(*sprintID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetBySprint(*sprintID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetBySprint(*sprintID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcSprint.Key, *sprintID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcSprint.Key, util.SvcSprint.Plural)
		link := util.AdminLink(util.SvcSprint.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, sprintID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminSprintDetail(sess, members, perms, estimates, standups, retros, actions, params, ctx, w))
	})
}
