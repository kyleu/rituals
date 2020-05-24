package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func TeamList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving teams"))
		}

		ctx.Title = util.KeyPluralTitle(util.SvcTeam.Key)
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcTeam.Key+".list"), util.SvcTeam.Key)
		return tmpl(templates.TeamList(sessions, ctx, w))
	})
}

func TeamNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(util.SvcTeam.Title, r.Form.Get("title"))
		sess, err := ctx.App.Team.New(title, ctx.Profile.UserID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating team session"))
		}
		return ctx.Route(util.SvcTeam.Key, util.KeyKey, sess.Slug), nil
	})
}

func TeamWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Team.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load team session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load team [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcTeam.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcTeam, ModelID: sess.ID, Slug: key, Title: sess.Title}
		permErrors, bc := check(&ctx, ctx.App.Team.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcTeam, permErrors, ctx, w)
		}

		ctx.Title = sess.Title

		return tmpl(templates.TeamWorkspace(sess, ctx, w))
	})
}
