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

func AdminStandupList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Daily Standup List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), util.SvcStandup.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		standups, err := ctx.App.Standup.List(params.Get("standup"))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminStandupList(standups, params, ctx, w))
	})
}

func AdminStandupDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		standupIDString := mux.Vars(r)["id"]
		standupID, err := uuid.FromString(standupIDString)
		if err != nil {
			return "", errors.New("invalid standup id [" + standupIDString + "]")
		}
		sess, err := ctx.App.Standup.GetByID(standupID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + standupIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.standup"), nil
		}

		params := paramSetFromRequest(r)

		members, err := ctx.App.Standup.Members.GetByModelID(standupID, params.Get("member"))
		if err != nil {
			return "", err
		}
		reports, err := ctx.App.Standup.GetReports(standupID, params.Get("report"))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcStandup.Key, standupID, params.Get("action"))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), util.SvcStandup.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup.detail", "id", standupIDString), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminStandupDetail(sess, members, reports, actions, params, ctx, w))
	})
}
