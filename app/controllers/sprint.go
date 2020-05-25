package controllers

import (
	"net/http"

	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/app/controllers/act"
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
			return "", errors.Wrap(err, "error retrieving sprints")
		}

		teams, err := ctx.App.Team.GetByMember(ctx.Profile.UserID, params.Get(util.SvcTeam.Key, ctx.Logger))
		if err != nil {
			return "", err
		}
		auths, err := ctx.App.Auth.GetByUserID(ctx.Profile.UserID, params.Get(util.KeyAuth, ctx.Logger))
		if err != nil {
			return "", err
		}

		ctx.Title = util.KeyPluralTitle(util.SvcSprint.Key)
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route(util.SvcSprint.Key+".list"), util.SvcSprint.Key)
		return tmpl(templates.SprintList(sessions, teams, auths, params.Get(util.SvcSprint.Key, ctx.Logger), ctx, w))
	})
}

func SprintNew(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		_ = r.ParseForm()
		title := util.ServiceTitle(util.SvcSprint, r.Form.Get("title"))
		startDate, err := util.FromYMD(r.Form.Get("startDate"))
		if err != nil {
			return "", err
		}
		endDate, err := util.FromYMD(r.Form.Get("endDate"))
		if err != nil {
			return "", err
		}
		teamID := getUUID(r.Form, util.SvcTeam.Key)
		perms := parsePerms(r.Form, teamID, nil)

		sess, err := ctx.App.Sprint.New(title, ctx.Profile.UserID, startDate, endDate, teamID)
		if err != nil {
			return "", errors.Wrap(err, "error creating sprint session")
		}

		_, err = ctx.App.Sprint.Permissions.SetAll(sess.ID, perms, ctx.Profile.UserID)
		if err != nil {
			return "", errors.Wrap(err, "error setting permissions for new session")
		}

		err = ctx.App.Socket.SendContentUpdate(util.SvcTeam.Key, teamID)
		if err != nil {
			return "", errors.Wrap(err, "cannot send content update")
		}
		return ctx.Route(util.SvcSprint.Key, util.KeyKey, sess.Slug), nil
	})
}

func SprintWorkspace(w http.ResponseWriter, r *http.Request) {
	act.Act(w, r, func(ctx web.RequestContext) (string, error) {
		key := mux.Vars(r)[util.KeyKey]
		sess, err := ctx.App.Sprint.GetBySlug(key)
		if err != nil {
			return "", errors.Wrap(err, "cannot load sprint session")
		}
		if sess == nil {
			ctx.Session.AddFlash("error:Can't load sprint [" + key + "]")
			act.SaveSession(w, r, ctx)
			return ctx.Route(util.SvcSprint.Key + ".list"), nil
		}

		params := PermissionParams{Svc: util.SvcSprint, ModelID: sess.ID, Slug: key, Title: sess.Title, TeamID: sess.TeamID}
		auths, permErrors, bc := check(&ctx, ctx.App.Sprint.Permissions, params)

		ctx.Breadcrumbs = bc

		if len(permErrors) > 0 {
			return permErrorTemplate(util.SvcSprint, permErrors, auths, ctx, w)
		}

		ctx.Title = sess.Title
		return tmpl(templates.SprintWorkspace(sess, auths, ctx, w))
	})
}
