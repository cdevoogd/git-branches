package checkout

import (
	"fmt"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:          "checkout",
	Short:        "Open a TUI for checking out branches",
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

	selection, err := tui.PromptForSelection(choices)
	if err != nil {
		return err
	}

	if selection == nil {
		fmt.Println("No branches were selected")
		return nil
	}

	err = git.CheckoutBranch(selection.Name)
	if err != nil {
		return err
	}

	return nil
}
