// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/clib"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/telemetry/httpmetrics"
	"github.com/kyleu/rituals/app/util"
)

func AppRoutes(as *app.State, logger util.Logger) fasthttp.RequestHandler {
	r := router.New()

	r.GET("/", controller.Home)
	r.GET("/healthcheck", clib.Healthcheck)
	r.GET("/about", clib.About)
	r.GET("/theme", clib.ThemeList)
	r.GET("/theme/{key}", clib.ThemeEdit)
	r.POST("/theme/{key}", clib.ThemeSave)
	r.GET(cutil.DefaultSearchPath, clib.Search)

	r.GET(cutil.DefaultProfilePath, clib.Profile)
	r.POST(cutil.DefaultProfilePath, clib.ProfileSave)
	r.GET("/auth/{key}", clib.AuthDetail)
	r.GET("/auth/callback/{key}", clib.AuthCallback)
	r.GET("/auth/logout/{key}", clib.AuthLogout)

	generatedRoutes(r)

	// $PF_SECTION_START(routes)$
	// $PF_SECTION_END(routes)$

	r.GET("/admin", clib.Admin)
	r.GET("/admin/sandbox", controller.SandboxList)
	r.GET("/admin/sandbox/{key}", controller.SandboxRun)
	r.GET("/admin/{path:*}", clib.Admin)

	r.GET("/favicon.ico", clib.Favicon)
	r.GET("/robots.txt", clib.RobotsTxt)
	r.GET("/assets/{_:*}", clib.Static)

	r.OPTIONS("/", controller.Options)
	r.OPTIONS("/{_:*}", controller.Options)
	r.NotFound = controller.NotFound

	clib.AppRoutesList = r.List()

	p := httpmetrics.NewMetrics(util.AppKey, logger)
	return fasthttp.CompressHandlerLevel(p.WrapHandler(r), fasthttp.CompressBestSpeed)
}
