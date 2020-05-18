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

func AdminRetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = "Retrospective List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), util.SvcRetro.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		retros, err := ctx.App.Retro.List(params.Get("retro"))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminRetroList(retros, params, ctx, w))
	})
}

func AdminRetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		retroIDString := mux.Vars(r)["id"]
		retroID, err := uuid.FromString(retroIDString)
		if err != nil {
			return "", errors.New("invalid retro id [" + retroIDString + "]")
		}
		sess, err := ctx.App.Retro.GetByID(retroID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + retroIDString + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.retro"), nil
		}

		params := paramSetFromRequest(r)

		members, err := ctx.App.Retro.Members.GetByModelID(retroID, params.Get("member"))
		if err != nil {
			return "", err
		}
		actions, err := ctx.App.Action.GetBySvcModel(util.SvcRetro.Key, retroID, params.Get("action"))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), util.SvcRetro.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro.detail", "id", retroIDString), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminRetroDetail(sess, members, actions, params, ctx, w))
	})
}
