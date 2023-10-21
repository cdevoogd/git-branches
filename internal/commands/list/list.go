package list

import (
	"fmt"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:          "list",
	Short:        "List the current repository's branches",
	SilenceUsage: true,
	RunE:         Execute,
}

func Execute(cmd *cobra.Command, args []string) error {
	branches, err := git.Branches()
	if err != nil {
		return fmt.Errorf("error loading branches: %w", err)
	}

	return tui.ListBranches(branches)
}
