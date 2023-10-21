package delete

import (
	"errors"
	"fmt"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:          "delete",
	Short:        "Open a TUI for deleting branches",
	SilenceUsage: true,
	RunE:         Execute,
}

func Execute(cmd *cobra.Command, args []string) error {
	branches, err := git.Branches()
	if err != nil {
		return fmt.Errorf("error loading branches: %w", err)
	}

	choices, err := tui.ConvertBranchesToChoices(branches)
	if err != nil {
		return fmt.Errorf("error converting branches to choices: %w", err)
	}

	selections, err := tui.PromptForSelections("Select branches to delete", choices)
	if err != nil {
		if errors.Is(err, tui.ErrQuit) {
			return nil
		}
		return err
	}

	if len(selections) == 0 {
		return nil
	}

	branchNames := make([]string, len(selections))
	for i, selection := range selections {
		branchNames[i] = selection.Name
	}

	err = git.DeleteBranches(branchNames)
	if err != nil {
		return err
	}

	return nil
}
