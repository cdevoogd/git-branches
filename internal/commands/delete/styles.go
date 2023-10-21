package delete

import (
	"github.com/cdevoogd/git-branches/internal/color"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().Foreground(color.White)
	selectedItemStyle = lipgloss.NewStyle().Foreground(color.Red)

	defaultPromptStyle      = lipgloss.NewStyle().Foreground(color.Red)
	finishPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(color.Green))
)

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
