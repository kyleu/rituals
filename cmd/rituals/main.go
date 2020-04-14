package main

import (
	"fmt"
	"os"

	"github.com/kyleu/rituals/internal/app/cli"
)

// Provisioned by ldflags
// nolint: gochecknoglobals
var (
	version    string
	commitHash string
)

func main() {
	Run()
}

func Run() {
	cmd := cli.Configure(version, commitHash)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
