package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/app/controllers/act"

	"github.com/kyleu/rituals.dev/app/team"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/sprint"
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
			return "", errors.WithStack(errors.Wrap(err, "error retrieving standups"))
		}

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcStandup.Key+".list"), util.SvcStandup.Key)
		return tmpl(templates.StandupList(sessions, ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(r.Form.Get("title"))
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		sprintID := getUUID(r.Form, util.SvcSprint.Key)
		sess, err := ctx.App.Standup.New(title, ctx.Profile.UserID, teamID, sprintID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "error creating standup session"))
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot send content update"))
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, sprintID)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot send content update"))
		}
		return ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Standup.GetBySlug(key)
		if err != nil {
			return "", errors.WithStack(errors.Wrap(err, "cannot load standup session"))
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load standup [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcStandup.Key + ".list"), nil
		}

		var tm *team.Session
		var tmTitle string
		if sess.TeamID != nil {
			tm, _ = ctx.App.Team.GetByID(*sess.TeamID)
			tmTitle = tm.Title
		}

		var spr *sprint.Session
		if sess.SprintID != nil {
			spr, _ = ctx.App.Sprint.GetByID(*sess.SprintID)
		}

		auths, currTeams, err := authsAndTeams(ctx, sess.TeamID)
		permErrors := ctx.App.Standup.Permissions.Check(util.SvcEstimate, sess.ID, auths, sess.TeamID, tmTitle, currTeams)
		if len(permErrors) > 0 {
			return permErrorTemplate(permErrors, ctx, w)
		}

		ctx.Title = sess.Title
		bc := web.BreadcrumbsSimple(ctx.Route(util.SvcStandup.Key+".list"), util.SvcStandup.Key)
		if spr != nil {
			bc = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key, util.KeyKey, spr.Slug), spr.Title)
		}
		bc = append(bc, web.BreadcrumbsSimple(ctx.Route(util.SvcStandup.Key, util.KeyKey, key), sess.Title)...)
		ctx.Breadcrumbs = bc

		return tmpl(templates.StandupWorkspace(sess, ctx, w))
	})
}
