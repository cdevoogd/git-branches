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

	defaultPromptStyle      = lipgloss.NewStyle().Foreground(color.Red)
	finishPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(color.Green))
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

// displayBranches returns a styled string that can be used to display the prompt to the user. It
// is meant to fulfil the prompt.Theme interface. This acts similar to the default theme but with
// colors that match the rest of the branch choice prompt.
func themePrompt(msg string, state prompt.State, model string) string {
	s := strings.Builder{}

	switch state {
	case prompt.StateNormal:
		s.WriteString(defaultPromptStyle.Render("?"))
	case prompt.StateFinish:
		s.WriteString(finishPromptPrefixStyle.Render("✔"))
	case prompt.StateError:
		s.WriteString(defaultPromptStyle.Render("✖"))
	}

	s.WriteString(" ")
	s.WriteString(msg)
	s.WriteString(" ")

	if state == prompt.StateNormal {
		s.WriteString(model)
	} else {
		s.WriteString(defaultPromptStyle.Render("…"))
		s.WriteString(" ")
		s.WriteString(model)
		s.WriteString("\n")
	}

	return s.String()
}
