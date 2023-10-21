package commands

import (
	"os"

	"github.com/cdevoogd/git-branches/internal/app"
	"github.com/cdevoogd/git-branches/internal/commands/delete"
	"github.com/cdevoogd/git-branches/internal/commands/list"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:     "git-branches",
	Short:   "A simple tool to replace Git's `git branch` command",
	Version: app.Version,
	RunE:    list.Execute,
}

func init() {
	RootCommand.AddCommand(list.Command)
	RootCommand.AddCommand(delete.Command)
}

func Execute() {
	err := RootCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}
