// Package controller - $PF_IGNORE$
package controller

import (
	"github.com/kyleu/rituals/app/workspace"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/views"
)

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		w, err := workspace.WorkspaceFromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Data = w
		return Render(rc, as, &views.Home{Teams: w.Teams, Sprints: w.Sprints, Estimates: w.Estimates, Standups: w.Standups, Retros: w.Retros}, ps)
	})
}
