package main // import github.com/kyleu/rituals

import (
	"context"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/cmd"
)

var (
	version = "2.3.1" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
