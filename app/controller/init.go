package controller

import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/util"
)

// Initialize app-specific system dependencies.
func initApp(_ *app.State, _ util.Logger) {
	user.SetPermissions(false,
		user.Perm("/admin", "github:kyleu.com", true),
		user.Perm("/admin", "github:rituals.dev", true),
		user.Perm("/admin", "*", false),
		user.Perm("/", "*", true),
	)
}

// Configure app-specific data for each request.
func initAppRequest(_ *app.State, _ *cutil.PageState) error {
	return nil
}

// Initialize system dependencies for the marketing site.
func initSite(_ *app.State, _ util.Logger) {
}

// Configure marketing site data for each request.
func initSiteRequest(_ *app.State, _ *cutil.PageState) error {
	return nil
}
