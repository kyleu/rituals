// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/rituals

import (
	"os"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/cmd"
	"github.com/kyleu/rituals/app/lib/log"
)

var (
	version = "2.0.11" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	logger, err := cmd.Run(&app.BuildInfo{Version: version, Commit: commit, Date: date})
	if err != nil {
		const msg = "exiting due to error"
		if logger == nil {
			println(log.Red.Add(err.Error())) //nolint:forbidigo
			println(log.Red.Add(msg))         //nolint:forbidigo
		} else {
			logger.Error(err)
			logger.Error(msg)
		}
		os.Exit(1)
	}
}
