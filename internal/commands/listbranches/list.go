package listbranches

import (
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/color"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
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

// Run will print out information (name, type, desc) about the given branches to stdout.
func Run(branches []*git.Branch) int {
	for _, branch := range branches {
		err := printBranch(branch)
		if err != nil {
			log.Error(err)
			return 1
		}
	}

	return 0
}

func printBranch(branch *git.Branch) error {
	nameStyle, ok := nameStyles[branch.Type]
	if !ok {
		return fmt.Errorf("no style is available for branch type %q", branch.Type)
	}

	s := strings.Builder{}
	s.WriteString(nameStyle.render(branch.Name))
	if branch.Description != "" {
		desc := descStyle.Render(fmt.Sprintf(" (%s)", branch.Description))
		s.WriteString(descStyle.Render(desc))
	}

	fmt.Println(s.String())
	return nil
}
