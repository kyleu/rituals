package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
)

func SprintCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateSprint(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, frm.Team, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New sprint created", model.PublicWebPath(), rc, ps)
	})
}

func SprintSave(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.sprint.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		_, msg, u, err := as.Services.Workspace.ActionSprint(ps.Context, slug, frm, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
