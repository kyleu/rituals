package site

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/lib/menu"
	"github.com/kyleu/rituals/app/lib/user"
	"github.com/kyleu/rituals/app/util"
)

const (
	keyAbout       = "about"
	keyContrib     = "contributing"
	keyCustomizing = "customizing"
	keyDownload    = "download"
	keyInstall     = "install"
	keyTech        = "technology"
)

func Menu(_ context.Context, _ *app.State, _ *user.Profile, _ user.Accounts, _ util.Logger) menu.Items {
	return menu.Items{
		{Key: keyInstall, Title: "Install", Icon: "code", Route: "/" + keyInstall},
		{Key: keyDownload, Title: "Download", Icon: "download", Route: "/" + keyDownload},
		{Key: keyCustomizing, Title: "Customizing", Icon: "code", Route: "/" + keyCustomizing},
		{Key: keyContrib, Title: "Contributing", Icon: "gift", Route: "/" + keyContrib},
		{Key: keyTech, Title: "Technology", Icon: "cog", Route: "/" + keyTech},
	}
}
