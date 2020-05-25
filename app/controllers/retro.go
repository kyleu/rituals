package controllers

import (
	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/query"
	"github.com/kyleu/rituals.dev/app/retro"
	"net/http"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Retro.GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))
		if err != nil {
			return "", errors.Wrap(err, "error retrieving retros")
		}

		teams, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		sprints, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = util.SvcRetro.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcRetro.Key+".list"), util.SvcRetro.Key)
		return tmpl(templates.RetroList(sessions, teams, sprints, auths, params.Get(util.SvcRetro.Key, ctx.Logger), ctx, w))
	})
}

func RetroNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(util.SvcRetro, r.Form.Get("title"))
		categoriesString := r.Form.Get("categories")
		categories := query.StringToArray(categoriesString)
		if len(categories) == 0 {
			categories = retro.DefaultCategories
		}
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		sprintID := getUUID(r.Form, util.SvcSprint.Key)
		perms := parsePerms(r.Form, teamID, sprintID)

		sess, err := ctx.App.Retro.New(title, ctx.Profile.UserID, categories, teamID, sprintID)
		if err != nil {
			return "", errors.Wrap(err, "error creating retro session")
		}

		_, err = ctx.App.Retro.Permissions.SetAll(sess.ID, perms, ctx.Profile.UserID)
		if err != nil {
			return "", errors.Wrap(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.Wrap(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, sprintID)
		if err != nil {
			return "", errors.Wrap(err, "cannot send content update")
		}

		return ctx.Route(util.SvcRetro.Key, util.KeyKey, sess.Slug), nil
	})
}

func RetroWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Retro.GetBySlug(key)
		if err != nil {
			return "", errors.Wrap(err, "cannot load retro session")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcRetro.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcRetro, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := check(&ctx, ctx.App.Retro.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcRetro, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.RetroWorkspace(sess, auths, ctx, w))
	})
}
