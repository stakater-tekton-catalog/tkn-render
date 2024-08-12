package main

import (
	"os"

	"github.com/stakater-tekton-catalog/tkn-render/internal/readme"
	"github.com/tektoncd/catlin/pkg/app"
	"github.com/tektoncd/catlin/pkg/cmd"
)

func main() {
	cli := app.New()
	catlin := cmd.Root(cli)
	catlin.AddCommand(readme.Command(cli))
	if err := catlin.Execute(); err != nil {
		os.Exit(1)
	}
}
