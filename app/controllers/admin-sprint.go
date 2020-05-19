package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminSprintList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Sprint List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint"), util.SvcSprint.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		sprints, err := ctx.App.Sprint.List(params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminSprintList(sprints, params, ctx, w))
	})
}

func AdminSprintDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		sprintIDString := mux.Vars(r)["id"]
		sprintID, err := uuid.FromString(sprintIDString)
		if err != nil {
			return "", errors.New("invalid sprint id [" + sprintIDString + "]")
		}
		sess, err := ctx.App.Sprint.GetByID(sprintID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + sprintIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.sprint"), nil
		}

		params := paramSetFromRequest(r)

		members, err := ctx.App.Sprint.Members.GetByModelID(sprintID, params.Get(util.KeyMember, ctx.Logger))
		if err != nil {
			return "", err
		}
		estimates, err := ctx.App.Estimate.GetBySprint(sprintID, params.Get(util.SvcEstimate.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		standups, err := ctx.App.Standup.GetBySprint(sprintID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		retros, err := ctx.App.Retro.GetBySprint(sprintID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcSprint.Key, sprintID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint"), util.SvcSprint.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.sprint.detail", "id", sprintIDString), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminSprintDetail(sess, members, estimates, standups, retros, actions, params, ctx, w))
	})
}
