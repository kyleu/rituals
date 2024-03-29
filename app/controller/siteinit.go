// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
)

// Initialize system dependencies for the marketing site.
func initSite(_ context.Context, _ *app.State, _ util.Logger) error {
	return nil
}

// Configure marketing site data for each request.
func initSiteRequest(*app.State, *cutil.PageState) error {
	return nil
}
