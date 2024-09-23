package main

import (
	"github.com/yjwj-3/content-release/cmd/api/app"
	"github.com/yjwj-3/content-release/cmd/cli"
	"os"
)

func main() {
	cmd := app.NewApiCommand()
	code := cli.Run(cmd)
	os.Exit(code)
}
