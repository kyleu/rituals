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

func RetroList(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		ctx.Title = util.SvcRetro.Title + " List"
		ctx.Breadcrumbs = adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		params := act.ParamSetFromRequest(r)
		retros, err := ctx.App.Retro.List(params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		return tmpl(templates.AdminRetroList(retros, params, ctx, w))
	})
}

func RetroDetail(w http.ResponseWriter, r *http.Request) {
	adminAct(w, r, func(ctx web.RequestContext) (string, error) {
		retroID := util.GetUUIDPointer(mux.Vars(r), util.KeyID)
		if retroID == nil {
			return "", errors.New("invalid retro id")
		}
		sess, err := ctx.App.Retro.GetByID(*retroID)
		if err != nil {
			return "", err
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + retroID.String() + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.AdminLink(util.SvcRetro.Key)), nil
		}

		params := act.ParamSetFromRequest(r)

		members := ctx.App.Retro.Members.GetByModelID(*retroID, params.Get(util.KeyMember, ctx.Logger))
		perms := ctx.App.Retro.Permissions.GetByModelID(*retroID, params.Get(util.KeyPermission, ctx.Logger))

		actions, err := ctx.App.Action.GetBySvcModel(util.SvcRetro.Key, *retroID, params.Get(util.KeyAction, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = sess.Title
		bc := adminBC(ctx, util.SvcRetro.Key, util.SvcRetro.Plural)
		link := util.AdminLink(util.SvcRetro.Key, util.KeyDetail)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(link, util.KeyID, retroID.String()), sess.Slug)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.AdminRetroDetail(sess, members, perms, actions, params, ctx, w))
	})
}
