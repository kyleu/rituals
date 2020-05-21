package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"
	"github.com/kyleu/rituals.dev/app/team"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals.dev/app/web"

	"github.com/kyleu/rituals.dev/gen/templates"
)

func SprintList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Sprint.GetByMember(ctx.Profile.UserID, params.Get(util.SvcSprint.Key, ctx.Logger))
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error retrieving sprints"))
		}

		ctx.Title = "Sprints"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key+".list"), util.SvcSprint.Key)
		return tmpl(templates.SprintList(sessions, ctx, w))
	})
}

func SprintNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		sess, err := ctx.App.Sprint.New(title, ctx.Profile.UserID, teamID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating sprint session"))
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot send content update"))
		}
		return ctx.Route(util.SvcSprint.Key, util.KeyKey, sess.Slug), nil
	})
}

func SprintWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Sprint.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load sprint session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcSprint.Key + ".list"), nil
		}

		var tm *team.Session
		var tmTitle string
		if sess.TeamID != nil {
			tm, _ = ctx.App.Team.GetByID(*sess.TeamID)
			tmTitle = tm.Title
		}

		auths, currTeams, err := authsAndTeams(ctx, sess.TeamID)
		permErrors := ctx.App.Sprint.Permissions.Check(util.SvcEstimate, sess.ID, auths, sess.TeamID, tmTitle, currTeams)
		if len(permErrors) > 0 {
			return permErrorTemplate(permErrors, ctx, w)
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key+".list"), util.SvcSprint.Key)
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.SprintWorkspace(sess, ctx, w))
	})
}
