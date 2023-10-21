package checkout

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
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

	handler := newHandler(branches)
	branch, err := handler.getBranchToCheckout()
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			return nil
		}
		return err
	}

	if len(branch) == 0 {
		fmt.Println("No branch was selected")
		return nil
	}

	err = git.CheckoutBranch(branch)
	if err != nil {
		return err
	}

	return nil
}

type handler struct {
	choices []choose.Choice
}

func newHandler(branches []*git.Branch) *handler {
	choices := make([]choose.Choice, len(branches))
	for i, branch := range branches {
		choices[i] = choose.Choice{
			Text: branch.Name,
			Note: getBranchInfo(branch),
		}
	}

	return &handler{choices: choices}
}

func (d *handler) getBranchToCheckout() (string, error) {
	msg := "Choose a branch to checkout:"
	return prompt.New().Ask(msg).AdvancedChoose(
		d.choices,
		choose.WithHelp(true),
	)
}

func getBranchInfo(branch *git.Branch) string {
	var str strings.Builder
	if branch.Type != git.BranchTypeNormal {
		str.WriteString(fmt.Sprintf("(%s)", branch.Type.String()))
	}
	if branch.Description != "" {
		str.WriteString(fmt.Sprintf(" %s", branch.Description))
	}
	return str.String()
}
