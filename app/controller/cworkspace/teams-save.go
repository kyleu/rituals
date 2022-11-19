package cworkspace

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
)

func TeamCreate(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := parseRequestForm(rc, ps.Profile.Name)
		if err != nil {
			return "", err
		}
		model, _, err := as.Services.Workspace.CreateTeam(ps.Context, frm.ID, frm.Title, ps.Profile.ID, frm.Name, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, "New team created", model.PublicWebPath(), rc, ps)
	})
}

func TeamSave(rc *fasthttp.RequestCtx) {
	controller.Act("workspace.team.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		slug, err := cutil.RCRequiredString(rc, "slug", false)
		if err != nil {
			return "", errors.Wrap(err, "must provide [slug] in path")
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}
		_, msg, u, err := as.Services.Workspace.ActionTeam(ps.Context, slug, frm, ps.Profile.ID, ps.Logger)
		if err != nil {
			return "", err
		}
		return controller.FlashAndRedir(true, msg, u, rc, ps)
	})
}
