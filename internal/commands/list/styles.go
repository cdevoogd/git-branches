package list

import (
	"fmt"

	"github.com/cdevoogd/git-branches/internal/color"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/charmbracelet/lipgloss"
)

type branchStyle struct {
	style  lipgloss.Style
	prefix string
}

func (b *branchStyle) render(branch string) string {
	return b.style.Render(fmt.Sprint(b.prefix, branch))
}

var (
	descStyle  = lipgloss.NewStyle().Foreground(color.White)
	nameStyles = map[git.BranchType]branchStyle{
		git.BranchTypeNormal:   {prefix: "  ", style: lipgloss.NewStyle().Bold(true).Foreground(color.White)},
		git.BranchTypeCurrent:  {prefix: "* ", style: lipgloss.NewStyle().Bold(true).Foreground(color.Green)},
		git.BranchTypeWorktree: {prefix: "+ ", style: lipgloss.NewStyle().Bold(true).Foreground(color.Cyan)},
	}
)
