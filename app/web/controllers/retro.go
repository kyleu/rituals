package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/model/retro"
	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func RetroList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		sessions := ctx.App.Retro.GetByMember(ctx.Profile.UserID, params.Get(util.SvcRetro.Key, ctx.Logger))
		teams := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcRetro.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcRetro.Key+".list"), util.SvcRetro.Plural)
		return tmpl(templates.RetroList(sessions, teams, sprints, auths, params.Get(util.SvcRetro.Key, ctx.Logger), ctx, w))
	})
}

func RetroNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()

		categoriesString := r.Form.Get(util.Plural(util.KeyCategory))
		categories := query.StringToArray(categoriesString)
		if len(categories) == 0 {
			categories = retro.DefaultCategories
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcRetro, r.Form, ctx.App.User)

		sess, err := ctx.App.Retro.New(sf.Title, ctx.Profile.UserID, sf.MemberName, categories, sf.TeamID, sf.SprintID)
		if err != nil {
			return eresp(err, "error creating retro session")
		}

		_, err = ctx.App.Retro.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam, sf.TeamID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint, sf.SprintID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcRetro.Key, util.KeyKey, sess.Slug), nil
	})
}

func RetroWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess := ctx.App.Retro.GetBySlug(key)
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load retro [" + key + "]")
			act.SaveSession(w, r, &ctx)
			return ctx.Route(util.SvcRetro.Key + ".list"), nil
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcRetro.Key, util.KeyKey, sess.Slug), nil
		}

		params := PermissionParams{Svc: util.SvcRetro, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := check(&ctx, ctx.App.Retro.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcRetro, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.RetroWorkspace(sess, auths, ctx, w))
	})
}
