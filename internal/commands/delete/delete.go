package delete

import (
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

	choices := make([]*tui.Choice, len(branches))
	for i, branch := range branches {
		choices[i], err = tui.NewChoiceFromBranch(branch)
		if err != nil {
			return fmt.Errorf("error converting branch to choice: %w", err)
		}
	}

	selections, err := tui.PromptForSelections(choices)
	if err != nil {
		return err
	}

	if len(selections) == 0 {
		fmt.Println("No branches were selected")
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
