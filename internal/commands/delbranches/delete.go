package delbranches

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/color"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/multichoose"
)

var (
	itemStyle         = lipgloss.NewStyle().Foreground(color.White)
	selectedItemStyle = lipgloss.NewStyle().Foreground(color.Red)
)

func Run(branches []*git.Branch) int {
	handler := newDeleteHandler(branches)
	toDelete, err := handler.getBranchesToDelete()
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Println("Exiting")
			return 0
		}

		log.Error(err)
		return 1
	}

	if len(toDelete) == 0 {
		fmt.Println("No branches were selected")
		return 0
	}

	err = git.DeleteBranches(toDelete)
	if err != nil {
		log.Error(err)
		return 1
	}

	return 0
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
