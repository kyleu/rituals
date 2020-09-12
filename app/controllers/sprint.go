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

func SprintList(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		params := npnweb.ParamSetFromRequest(r)

		sessions := app.Svc(ctx.App).Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		teams := app.Svc(ctx.App).Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		auths := ctx.App.Auth().GetByUserID(ctx.Profile.UserID, params.Get(npncore.KeyAuth, ctx.Logger))

		ctx.Title = npncore.PluralTitle(util.SvcSprint.Key)
		ctx.Breadcrumbs = npnweb.Breadcrumbs{npnweb.BreadcrumbSelf(util.SvcSprint.Plural)}
		return npncontroller.T(templates.SprintList(sessions, teams, auths, params.Get(util.SvcSprint.Key, ctx.Logger), ctx, w))
	})
}

func SprintNew(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		_ = r.ParseForm()

		startDate, err := npncore.FromYMD(r.Form.Get("startDate"))
		if err != nil {
			return npncontroller.EResp(err)
		}
		endDate, err := npncore.FromYMD(r.Form.Get("endDate"))
		if err != nil {
			return npncontroller.EResp(err)
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcSprint, r.Form, ctx.App.User())

		sess, err := app.Svc(ctx.App).Sprint.New(sf.Title, ctx.Profile.UserID, sf.MemberName, startDate, endDate, sf.TeamID)
		if err != nil {
			return npncontroller.EResp(err, "error creating sprint session")
		}

		_, err = app.Svc(ctx.App).Sprint.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return npncontroller.EResp(err, "error setting permissions for new session")
		}

		err = socket.SendContentUpdate(app.Svc(ctx.App).Socket, util.SvcTeam.Key, sf.TeamID)
		if err != nil {
			return npncontroller.EResp(err, "cannot send content update")
		}
		return ctx.Route(util.SvcSprint.Key, npncore.KeyKey, sess.Slug), nil
	})
}

func SprintWorkspace(w http.ResponseWriter, r *http.Request) {
	npncontroller.Act(w, r, func(ctx *npnweb.RequestContext) (string, error) {
		key := mux.Vars(r)[npncore.KeyKey]
		sess := app.Svc(ctx.App).Sprint.GetBySlug(key)
		if sess == nil {
			msg := "can't load sprint [" + key + "]"
			return npncontroller.FlashAndRedir(false, msg, util.SvcSprint.Key+".list", w, r, ctx)
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcSprint.Key, npncore.KeyKey, sess.Slug), nil
		}

		params := &PermissionParams{Svc: util.SvcSprint, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID}
		auths, permErrors, bc := CheckPerms(ctx, app.Svc(ctx.App).Sprint.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return PermErrorTemplate(util.SvcSprint, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return npncontroller.T(templates.SprintWorkspace(sess, auths, ctx, w))
	})
}

func SprintExport(w http.ResponseWriter, r *http.Request) {
	f := func(key string, ctx *npnweb.RequestContext) ExportParams {
		sess := app.Svc(ctx.App).Sprint.GetBySlug(key)
		if sess == nil {
			return ExportParams{}
		}
		return ExportParams{
			ModelID: &sess.ID,
			Slug:    sess.Slug,
			Title:   sess.Title,
			Path:    ctx.Route(util.SvcSprint.Key, npncore.KeyKey, sess.Slug),
			PermSvc: app.Svc(ctx.App).Sprint.Data.Permissions,
		}
	}
	ExportAct(util.SvcSprint, f, w, r)
}
