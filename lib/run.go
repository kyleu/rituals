package lib

import (
	"emperror.dev/errors"
	"github.com/kyleu/rituals.dev/internal/app/cli"
)

func Run() {
	ai, err := cli.InitApp("0.0.0", "master")
	if err != nil {
		panic(errors.WithStack(err))
	}
	err = cli.MakeServer(ai, "127.0.0.1", 6660)
	if err != nil {
		panic(errors.WithStack(err))
	}
}
