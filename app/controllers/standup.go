package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/socket"

	"github.com/kyleu/npn/npncontroller"
	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func StandupList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Svc(ctx.App).Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		teams := app.Svc(ctx.App).Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		sprints := app.Svc(ctx.App).Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcStandup.Plural)}
		return npncontroller.T(templates.StandupList(sessions, teams, sprints, auths, params.Get(util.SvcStandup.Key, ctx.Logger), ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcStandup, r.Form, ctx.App.User())

		sess, err := app.Svc(ctx.App).Standup.New(sf.Title, ctx.Profile.UserID, sf.MemberName, sf.TeamID, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "error creating standup session")
		}

		_, err = app.Svc(ctx.App).Standup.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return npncontroller.EResp(err, "error setting permissions for new session")
		}

		err = socket.SendContentUpdate(app.Svc(ctx.App).Socket, util.SvcTeam.Key, sf.TeamID)
		if err != nil {
			return npncontroller.EResp(err, "cannot send content update")
		}
		err = socket.SendContentUpdate(app.Svc(ctx.App).Socket, util.SvcSprint.Key, sf.SprintID)
		if err != nil {
			return npncontroller.EResp(err, "cannot send content update")
		}

		return ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Svc(ctx.App).Standup.GetBySlug(key)
		if sess == nil {
			msg := "can't load standup [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcStandup.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &PermissionParams{Svc: util.SvcStandup, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID, SprintID: sess.SprintID}
		auths, permErrors, bc := CheckPerms(ctx, app.Svc(ctx.App).Standup.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return PermErrorTemplate(util.SvcStandup, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return npncontroller.T(templates.StandupWorkspace(sess, auths, ctx, w))
	})
}

func StandupExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) ExportParams {
		sess := app.Svc(ctx.App).Standup.GetBySlug(key)
		if sess == nil {
			return ExportParams{}
		}
		return ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcStandup.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Svc(ctx.App).Standup.Data.Permissions,
		}
	}
	ExportAct(util.SvcStandup, f, w, r)
}
