package style

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	NormalItem      = Style{lipgloss.NewStyle()}
	HighlightedItem = Style{lipgloss.NewStyle().Foreground(lipgloss.Color("10"))}
	Note            = Style{lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	})}
)

type Style struct {
	lipgloss.Style
}

func NewStyle() Style {
	return Style{Style: lipgloss.NewStyle()}
}

func (s Style) Renderf(format string, a ...any) string {
	return s.Render(fmt.Sprintf(format, a...))
}

func (s Style) Writef(b *strings.Builder, format string, a ...any) {
	b.WriteString(s.Renderf(format, a...))
}
