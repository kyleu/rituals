package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"

	"emperror.dev/errors"
	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func AdminRetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin."+util.SvcRetro.Key), util.SvcRetro.Key)...)
		ctx.Breadcrumbs = bc

		params := paramSetFromRequest(r)
		retros, err := ctx.App.Retro.List(params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminRetroList(retros, params, ctx, w))
	})
}

func AdminRetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		retroID := getUUIDPointer(mux.Vars(r), "id")
		if retroID == nil {
			return "", errors.New("invalid retro id")
		}
		sess, err := ctx.App.Retro.GetByID(*retroID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + retroID.String() + "]")
			saveSession(w, r, ctx)
			return ctx.Route("admin.retro"), nil
		}

		params := paramSetFromRequest(r)

		members := ctx.App.Retro.Members.GetByModelID(*retroID, params.Get(util.KeyMember, ctx.Logger))

		actions, err := ctx.App.Action.GetBySvcModel(util.SvcRetro.Key, *retroID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route("admin"), "admin")
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro"), util.SvcRetro.Key)...)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route("admin.retro.detail", "id", retroID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminRetroDetail(sess, members, actions, params, ctx, w))
	})
}
