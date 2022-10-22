// Package controller - $PF_IGNORE$
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views"
)

var homeContent = util.ValueMap{
	"_": util.AppName,
	"urls": map[string]string{
		"TODO": "/todo",
	},
}

func Home(rc *fasthttp.RequestCtx) {
	Act("home", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Data = homeContent
		return Render(rc, as, &views.Home{}, ps)
	})
}
