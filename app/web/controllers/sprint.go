package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web/act"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)

		sessions := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		teams := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		auths := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))

		ctx.Title = util.PluralTitle(util.SvcSprint.Key)
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key+".list"), util.SvcSprint.Plural)
		return tmpl(templates.SprintList(sessions, teams, auths, params.Get(util.SvcSprint.Key, ctx.Logger), ctx, w))
	})
}

func SprintNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()

		startDate, err := util.FromYMD(r.Form.Get("startDate"))
		if err != nil {
			return eresp(err, "")
		}
		endDate, err := util.FromYMD(r.Form.Get("endDate"))
		if err != nil {
			return eresp(err, "")
		}

		sf := parseSessionForm(ctx.Profile.UserID, util.SvcSprint, r.Form, ctx.App.User)

		sess, err := ctx.App.Sprint.New(sf.Title, ctx.Profile.UserID, startDate, endDate, sf.TeamID)
		if err != nil {
			return eresp(err, "error creating sprint session")
		}

		_, err = ctx.App.Sprint.Data.Permissions.SetAll(sess.ID, sf.Perms, ctx.Profile.UserID)
		if err != nil {
			return eresp(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam, sf.TeamID)
		if err != nil {
			return eresp(err, "cannot send content update")
		}
		return ctx.Route(util.SvcSprint.Key, util.KeyKey, sess.Slug), nil
	})
}

func SprintWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess := ctx.App.Sprint.GetBySlug(key)
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcSprint.Key + ".list"), nil
		}
		if sess.Slug != key {
			return ctx.Route(util.SvcSprint.Key, util.KeyKey, sess.Slug), nil
		}

		params := PermissionParams{Svc: util.SvcSprint, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID}
		auths, permErrors, bc := check(&ctx, ctx.App.Sprint.Data.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcSprint, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.SprintWorkspace(sess, auths, ctx, w))
	})
}
