package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	normalItemStyle      = style{lipgloss.NewStyle()}
	highlightedItemStyle = style{lipgloss.NewStyle().Foreground(lipgloss.Color("10"))}
	noteStyle            = style{lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	})}
)

type style struct {
	lipgloss.Style
}

func (s style) Write(b *strings.Builder, format string, a ...any) {
	b.WriteString(s.Render(fmt.Sprintf(format, a...)))
}
