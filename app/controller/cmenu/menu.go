package cmenu

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/lib/filter"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/lib/sandbox"
	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/util"
)

func MenuFor(
	ctx context.Context, as *app.State, isAuthed bool, isAdmin bool, profile *user.Profile, params filter.ParamSet, logger util.Logger,
) (menu.Items, any, error) {
	ctx, sp, _ := telemetry.StartSpan(ctx, "menu", logger)
	defer sp.Complete()
	var ret menu.Items
	var data any
	// $PF_SECTION_START(menu)$
	ws, data, err := workspaceMenu(ctx, as, params, profile, logger)
	if err != nil {
		return nil, nil, err
	}

	ret = append(ret, ws...)
	if isAdmin {
		ret = append(ret, menu.Separator)
	}
	if isAdmin {
		ret = append(ret, generatedMenu()...)
	}
	if isAdmin {
		admin := &menu.Item{Key: "admin", Title: "Settings", Description: "System-wide settings and preferences", Icon: "cog", Route: "/admin"}
		ret = append(ret, menu.Separator, sandbox.Menu(ctx), menu.Separator, admin)
	}
	const aboutDesc = "Get assistance and advice for using " + util.AppName
	ret = append(ret, menu.Separator, &menu.Item{Key: "about", Title: "About", Description: aboutDesc, Icon: "question", Route: "/about"})
	// $PF_SECTION_END(menu)$
	return ret, data, nil
}
