package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminStandupList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcStandup.Title + " List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), util.SvcStandup.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		standups, err := ctx.App.Standup.List(params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminStandupList(standups, params, ctx, w))
	})
}

func AdminStandupDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		standupID := getUUIDPointer(mux.Vars(r), "id")
		if standupID == nil {
			return "", errors.New("invalid standup id")
		}
		sess, err := ctx.App.Standup.GetByID(*standupID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + standupID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.standup"), nil
		}

		params := paramSetFromRequest(r)

		members := ctx.App.Standup.Members.GetByModelID(*standupID, params.Get(util.KeyMember, ctx.Logger))

		reports, err := ctx.App.Standup.GetReports(*standupID, params.Get(util.KeyReport, ctx.Logger))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcStandup.Key, *standupID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup"), util.SvcStandup.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.standup.detail", "id", standupID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminStandupDetail(sess, members, reports, actions, params, ctx, w))
	})
}
