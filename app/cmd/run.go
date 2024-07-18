package cmd

import (
	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return util.RootLogger, err
	}
	return util.RootLogger, nil
}
