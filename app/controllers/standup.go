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

func StandupList(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		params := act.ParamSetFromRequest(r)
		sessions, err := ctx.App.Standup.GetByMember(ctx.Profile.UserID, params.Get(util.SvcStandup.Key, ctx.Logger))
		if err != nil {
			return "", errors.Wrap(err, "error retrieving standups")
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

		ctx.Title = util.SvcStandup.PluralTitle
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcStandup.Key+".list"), util.SvcStandup.Key)
		return tmpl(templates.StandupList(sessions, teams, sprints, auths, params.Get(util.SvcStandup.Key, ctx.Logger), ctx, w))
	})
}

func StandupNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(util.SvcStandup, r.Form.Get("title"))
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		sprintID := getUUID(r.Form, util.SvcSprint.Key)
		sess, err := ctx.App.Standup.New(title, ctx.Profile.UserID, teamID, sprintID)
		if err != nil {
			return "", errors.Wrap(err, "error creating standup session")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.Wrap(err, "cannot send content update")
		}
		err = ctx.App.Socket.SendContentUpdate(util.SvcSprint.Key, sprintID)
		if err != nil {
			return "", errors.Wrap(err, "cannot send content update")
		}
		return ctx.Route(util.SvcStandup.Key, util.KeyKey, sess.Slug), nil
	})
}

func StandupWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Standup.GetBySlug(key)
		if err != nil {
			return "", errors.Wrap(err, "cannot load standup session")
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
