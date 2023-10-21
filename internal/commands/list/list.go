package list

import (
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List the current repository's branches",
	RunE:  Execute,
}

func Execute(cmd *cobra.Command, args []string) error {
	branches, err := git.Branches()
	if err != nil {
		return fmt.Errorf("error loading branches: %w", err)
	}

	for _, branch := range branches {
		err := printBranch(branch)
		if err != nil {
			return fmt.Errorf("error printing branch: %w", err)
		}
	}

	return nil
}

func printBranch(branch *git.Branch) error {
	nameStyle, ok := nameStyles[branch.Type]
	if !ok {
		return fmt.Errorf("no style is available for branch type %q", branch.Type)
	}

	s := strings.Builder{}
	s.WriteString(nameStyle.render(branch.Name))
	if branch.Description != "" {
		desc := descStyle.Render(fmt.Sprintf(" (%s)", branch.Description))
		s.WriteString(descStyle.Render(desc))
	}

	fmt.Println(s.String())
	return nil
}
