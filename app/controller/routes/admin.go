// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/rituals/app/controller/clib"
)

func adminRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/admin", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/database", clib.DatabaseList)
	makeRoute(r, http.MethodGet, "/admin/database/{key}", clib.DatabaseDetail)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/{act}", clib.DatabaseAction)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/tables/{schema}/{table}", clib.DatabaseTableView)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/tables/{schema}/{table}/stats", clib.DatabaseTableStats)
	makeRoute(r, http.MethodPost, "/admin/database/{key}/sql", clib.DatabaseSQLRun)
	makeRoute(r, http.MethodGet, "/admin/sandbox", clib.SandboxList)
	makeRoute(r, http.MethodGet, "/admin/sandbox/{key}", clib.SandboxRun)
	makeRoute(r, http.MethodGet, "/admin/{path:.*}", clib.Admin)
	makeRoute(r, http.MethodPost, "/admin/{path:.*}", clib.Admin)
}
