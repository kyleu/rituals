package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		sessions := ctx.App.Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		teams := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = web.Breadcrumbs{web.BreadcrumbSelf(util.SvcStandup.Plural)}
		return act.T(templates.StandupList(sessions, teams, sprints, auths, params.Get(util.SvcStandup.Key, ctx.Logger), ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		_ = r.ParseForm()

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcStandup, r.Form, ctx.App.User)

		sess, err := ctx.App.Standup.New(sf.Title, ctx.Profile.UserID, sf.MemberName, sf.TeamID, sf.SprintID)
		if err != nil {
			return act.EResp(err, "error creating standup session")
		}

		_, err = ctx.App.Standup.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return act.EResp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam, sf.TeamID)
		if err != nil {
			return act.EResp(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint, sf.SprintID)
		if err != nil {
			return act.EResp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx *web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess := ctx.App.Standup.GetBySlug(key)
		if sess == nil {
			msg := "can't load standup [" + key + "]"
			return act.FlashAndRedir(false, msg, util.SvcStandup.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug), nil
		}

		params := &act.PermissionParams{Svc: util.SvcStandup, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := act.CheckPerms(ctx, ctx.App.Standup.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return act.PermErrorTemplate(util.SvcStandup, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return act.T(templates.StandupWorkspace(sess, auths, ctx, w))
	})
}

func StandupExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *web.RequestContext) act.ExportParams {
		sess := ctx.App.Standup.GetBySlug(key)
		if sess == nil {
			return act.ExportParams{}
		}
		return act.ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug),
			PermSvc: ctx.App.Standup.Data.Permissions,
		}
	}
	act.ExportAct(util.SvcStandup, f, w, r)
}
