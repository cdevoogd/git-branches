package tui

import (
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui/style"
)

func ListBranches(branches []*git.Branch) error {
	builder := &strings.Builder{}
	for _, branch := range branches {
		addBranch(builder, branch)
	}

	_, err := fmt.Print(builder.String())
	return err
}

func addBranch(s *strings.Builder, branch *git.Branch) {
	symbol, branchStyle := getBranchStyles(branch)

	s.WriteString(symbol)
	s.WriteString(" ")
	s.WriteString(branchStyle.Render(branch.Name))
	if branch.Description != "" {
		s.WriteString(" ")
		s.WriteString(style.Note.Render(branch.Description))
	}
	s.WriteString("\n")
}

func getBranchStyles(branch *git.Branch) (symbol string, styler style.Style) {
	switch branch.Type {
	case git.BranchTypeCurrent:
		return "*", style.CurrentBranch
	case git.BranchTypeWorktree:
		return "+", style.WorktreeBranch
	default:
		return " ", style.NormalBranch
	}
}
