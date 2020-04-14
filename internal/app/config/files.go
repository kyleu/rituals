package config

import (
	"github.com/kyleu/rituals.dev/internal/app/util"
	"path/filepath"

	"github.com/kirsle/configdir"
)

var cfgDir = ""

func ConfigPath(filename string) string {
	if cfgDir == "" {
		cfgDir = configdir.LocalConfig(util.AppName)
		_ = configdir.MakePath(cfgDir)
	}
	return filepath.Join(cfgDir, filename)
}
