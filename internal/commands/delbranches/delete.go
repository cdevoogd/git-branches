package delbranches

import (
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/multichoose"
)

var (
	itemStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F14E32"))
)

func Run(branches []*git.Branch) error {
	handler := newDeleteHandler(branches)
	toDelete, err := handler.getBranchesToDelete()
	if err != nil {
		return err
	}

	return git.DeleteBranches(toDelete)
}

type deleteHandler struct {
	branches map[string]*git.Branch
}

func newDeleteHandler(branches []*git.Branch) *deleteHandler {
	handler := &deleteHandler{
		branches: make(map[string]*git.Branch),
	}

	for _, branch := range branches {
		handler.branches[branch.Name] = branch
	}

	return handler
}

func (d *deleteHandler) getBranchesToDelete() ([]string, error) {
	var choices []string
	for name := range d.branches {
		choices = append(choices, name)
	}

	msg := "Choose branches to delete:"
	return prompt.New().Ask(msg).MultiChoose(
		choices,
		multichoose.WithTheme(d.displayBranches),
		multichoose.WithHelp(true),
	)
}

// displayBranches returns a styled string that can be used to display the prompt to the user. It
// is meant to fulfil the multichoose.Theme interface.
func (d *deleteHandler) displayBranches(branches []string, cursor int, isSelected multichoose.IsSelected) string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i, name := range branches {
		symbol, style := getDisplayStyle(i == cursor, isSelected(i))
		s.WriteString(style.Render(fmt.Sprintf("[%s] %s", symbol, name)))

		branch, ok := d.branches[name]
		if ok && branch.Description != "" {
			s.WriteString(style.Render(fmt.Sprintf(" (%s)", branch.Description)))
		}

		s.WriteString("\n")
	}

	return s.String()
}

func getDisplayStyle(itemAtCursor, selected bool) (symbol string, style lipgloss.Style) {
	if itemAtCursor {
		if selected {
			return "x", selectedItemStyle
		}
		return "•", selectedItemStyle
	}
	if selected {
		return "x", itemStyle
	}
	return " ", itemStyle
}
