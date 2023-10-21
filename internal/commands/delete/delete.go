package delete

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/multichoose"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "delete",
	Short: "Open a TUI for deleting branches",
	RunE:  Execute,
}

func Execute(cmd *cobra.Command, args []string) error {
	branches, err := git.Branches()
	if err != nil {
		return fmt.Errorf("error loading branches: %w", err)
	}

	handler := newDeleteHandler(branches)
	toDelete, err := handler.getBranchesToDelete()
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			return nil
		}
		return err
	}

	if len(toDelete) == 0 {
		fmt.Println("No branches were selected")
		return nil
	}

	err = git.DeleteBranches(toDelete)
	if err != nil {
		return err
	}

	return nil
}

type deleteHandler struct {
	choices  []string
	branches map[string]*git.Branch
}

func newDeleteHandler(branches []*git.Branch) *deleteHandler {
	handler := &deleteHandler{
		choices:  make([]string, len(branches)),
		branches: make(map[string]*git.Branch),
	}

	for i, branch := range branches {
		handler.choices[i] = branch.Name
		handler.branches[branch.Name] = branch
	}

	return handler
}

func (d *deleteHandler) getBranchesToDelete() ([]string, error) {
	msg := "Choose branches to delete:"
	return prompt.New(prompt.WithTheme(themePrompt)).Ask(msg).MultiChoose(
		d.choices,
		multichoose.WithTheme(d.themeChoices),
		multichoose.WithHelp(true),
	)
}

// themeChoices returns a styled string that can be used to display the branch choices to the user.
// It is meant to fulfil the multichoose.Theme interface.
func (d *deleteHandler) themeChoices(branches []string, cursor int, isSelected multichoose.IsSelected) string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i, name := range branches {
		symbol, style := getDisplayStyle(i == cursor, isSelected(i))
		s.WriteString(style.Render(fmt.Sprintf("[%s] %s", symbol, name)))

		branch, ok := d.branches[name]
		if ok {
			if branch.Type != git.BranchTypeNormal {
				s.WriteString(style.Render(fmt.Sprintf(" (%s)", branch.Type.String())))
			}

			if branch.Description != "" {
				s.WriteString(style.Render(fmt.Sprintf(" (%s)", branch.Description)))
			}
		}

		s.WriteString("\n")
	}

	return s.String()
}
