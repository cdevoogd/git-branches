package tui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui/style"
)

type Choice struct {
	Name string
	Note string
}

func NewChoiceFromBranch(branch *git.Branch) (*Choice, error) {
	if branch == nil {
		return nil, errors.New("branch is nil")
	}

	var note strings.Builder
	if branch.Type != git.BranchTypeNormal {
		note.WriteString(fmt.Sprintf("(%s) ", branch.Type.String()))
	}
	if branch.Description != "" {
		note.WriteString(branch.Description)
	}

	return &Choice{
		Name: branch.Name,
		Note: strings.TrimSpace(note.String()),
	}, nil
}

type choiceRenderContext struct {
	hovered  bool
	selected bool

	normalPrefix   string
	hoveredPrefix  string
	selectedPrefix string
}

func (c Choice) render(ctx *choiceRenderContext) string {
	prefix := ctx.normalPrefix
	styler := style.NormalItem
	if ctx.hovered {
		prefix = ctx.hoveredPrefix
		styler = style.HighlightedItem
	}
	if ctx.selected {
		prefix = ctx.selectedPrefix
	}

	str := &strings.Builder{}

	styler.Writef(str, "%s%s", prefix, c.Name)
	if c.Note != "" {
		style.Note.Writef(str, "  %s", c.Note)
	}

	return str.String()
}
