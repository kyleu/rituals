package lib

import (
	"github.com/kyleu/rituals.dev/app/cli"
)

func Run() {
	ai, err := cli.InitApp("0.0.0", "master")
	if err != nil {
		panic(err)
	}
	err = cli.MakeServer(ai, "127.0.0.1", 6660)
	if err != nil {
		panic(err)
	}
}
