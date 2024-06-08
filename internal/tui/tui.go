package tui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	ErrQuit = errors.New("user exited prompt")
)

func ListBranches(branches []*git.Branch) error {
	out := &strings.Builder{}
	for _, branch := range branches {
		switch branch.Type {
		case git.BranchTypeCurrent:
			out.WriteString(listCurrentBranchStyle.Render(branch.Name))
		case git.BranchTypeWorktree:
			out.WriteString(listWorktreeBranchStyle.Render(branch.Name))
		default:
			out.WriteString(listNormalBranchStyle.Render(branch.Name))
		}

		if branch.Description != "" {
			out.WriteString(descriptionStyle.Render(branch.Description))
		}

		out.WriteRune('\n')
	}

	_, err := fmt.Print(out.String())
	return err
}

func RunSingleSelect(items []*Item) (*Item, error) {
	model := NewSingleSelectModel(items)
	_, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	return model.SelectedItem(), model.Error()
}

func RunMultiSelect(items []*Item) ([]*Item, error) {
	model := NewMultiSelectModel(items)
	_, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	return model.SelectedItems(), model.Error()
}
