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

	items, err := tui.ItemsFromBranches(branches)
	if err != nil {
		return fmt.Errorf("error converting branches to items: %w", err)
	}

	selections, err := tui.RunMultiSelect(items)
	if err != nil {
		if errors.Is(err, tui.ErrQuit) {
			fmt.Println("The prompt was manually exited")
			return nil
		}
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

	fmt.Printf("Deleting %d branches\n", len(selections))
	err = git.DeleteBranches(branchNames)
	if err != nil {
		return err
	}

	return nil
}
