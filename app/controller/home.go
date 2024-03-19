package controller

import (
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/workspace"
	"github.com/kyleu/rituals/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Act("home", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ws, err := workspace.FromAny(ps.Data)
		if err != nil {
			return "", err
		}
		ps.Data = ws
		return Render(w, r, as, &views.Home{Teams: ws.Teams, Sprints: ws.Sprints, Estimates: ws.Estimates, Standups: ws.Standups, Retros: ws.Retros}, ps)
	})
}
