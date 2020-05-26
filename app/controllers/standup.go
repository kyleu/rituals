package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "error retrieving standups")
		}

		teams, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}
		sprints, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return eresp(err, "")
		}

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcStandup.Key+".list"), util.SvcStandup.Key)
		return tmpl(templates.StandupList(sessions, teams, sprints, auths, params.Get(util.SvcStandup.Key, ctx.Logger), ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()

		r, err := parseSessionForm(ctx.Profile.UserID, util.SvcStandup, r.Form, ctx.App.User)
		if err != nil {
			return eresp(err, "cannot parse form")
		}

		sess, err := ctx.App.Standup.New(r.Title, ctx.Profile.UserID, r.TeamID, r.SprintID)
		if err != nil {
			return eresp(err, "error creating standup session")
		}

		_, err = ctx.App.Standup.Permissions.SetAll(sess.ID, r.Perms, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, r.TeamID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, r.SprintID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Standup.GetBySlug(key)
		if err != nil {
			return eresp(err, "cannot load standup session")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcStandup.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcStandup, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := check(&ctx, ctx.App.Standup.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcStandup, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.StandupWorkspace(sess, auths, ctx, w))
	})
}
