package checkout

import (
	"errors"
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

	choices, err := tui.ConvertBranchesToChoices(branches)
	if err != nil {
		return fmt.Errorf("error converting branches to choices: %w", err)
	}

	selection, err := tui.PromptForSelection("Select a branch to checkout", choices)
	if err != nil {
		if errors.Is(err, tui.ErrQuit) {
			return nil
		}
		return err
	}

	if selection == nil {
		return nil
	}

	err = git.CheckoutBranch(selection.Name)
	if err != nil {
		return err
	}

	return nil
}
