// Content managed by Project Forge, see [projectforge.md] for details.
package cmenu

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/lib/sandbox"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/util"
)

func MenuFor(ctx context.Context, isAuthed bool, isAdmin bool, as *app.State, logger util.Logger) (menu.Items, error) {
	ctx, span, logger := telemetry.StartSpan(ctx, "menu:generate", logger)
	defer span.Complete()
	_ = logger

	var ret menu.Items
	// $PF_SECTION_START(routes_start)$
	ws, err := workspaceMenu(ctx, as, logger)
	if err != nil {
		return nil, err
	}

	ret = append(ret, ws...)
	if isAdmin {
		ret = append(ret, menu.Separator)
	}
	// $PF_SECTION_END(routes_start)$
	if isAdmin {
		ret = append(ret, generatedMenu()...)
	}
	// $PF_SECTION_START(routes_end)$
	if isAdmin {
		admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
		ret = append(ret, menu.Separator, graphQLMenu(ctx, as.GraphQL), menu.Separator, sandbox.Menu(ctx), menu.Separator, admin)
	}
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(routes_end)$
	return ret, nil
}
