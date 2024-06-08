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

	items, err := tui.ItemsFromBranches(branches)
	if err != nil {
		return fmt.Errorf("error converting branches to items: %w", err)
	}

	selection, err := tui.RunSingleSelect(items)
	if err != nil {
		if errors.Is(err, tui.ErrQuit) {
			fmt.Println("The prompt was manually exited")
			return nil
		}
		return err
	}

	if selection == nil {
		fmt.Println("No branch was selected")
		return nil
	}

	fmt.Println("Checking out:", selection.Name)
	err = git.CheckoutBranch(selection.Name)
	if err != nil {
		return err
	}

	return nil
}
