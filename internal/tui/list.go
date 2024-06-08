package tui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Name     string
	Note     string
	Selected bool
}

func NewItemFromBranch(branch *git.Branch) (*Item, error) {
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

	return &Item{
		Name: branch.Name,
		Note: strings.TrimSpace(note.String()),
	}, nil
}

func ItemsFromBranches(branches []*git.Branch) ([]*Item, error) {
	choices := make([]*Item, len(branches))
	for i, branch := range branches {
		choice, err := NewItemFromBranch(branch)
		if err != nil {
			return nil, err
		}
		choices[i] = choice
	}
	return choices, nil
}

type RenderItemFunc func(item *Item, hovered bool) string

type listModel struct {
	items      []*Item
	cursor     int
	renderItem RenderItemFunc
	paginator  paginator.Model
}

func newListModel(items []*Item, renderItem RenderItemFunc) *listModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = activePaginationDotStyle
	p.InactiveDot = inactivePaginationDotStyle
	p.PerPage = 10
	p.SetTotalPages(len(items))

	return &listModel{
		items:      items,
		renderItem: renderItem,
		paginator:  p,
	}
}

func (m *listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	var cmd tea.Cmd
	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m listModel) View() string {
	view := &strings.Builder{}
	for i, item := range m.CurrentPageItems() {
		hovered := m.cursor == i
		view.WriteString(m.renderItem(item, hovered))
		view.WriteRune('\n')
	}

	if m.paginator.TotalPages > 1 {
		view.WriteString(m.paginator.View())
	}
	return view.String()
}

func (m *listModel) SelectedItems() []*Item {
	var selected []*Item
	for _, c := range m.items {
		if c.Selected {
			selected = append(selected, c)
		}
	}
	return selected
}

func (m *listModel) CurrentPageItems() []*Item {
	start, end := m.paginator.GetSliceBounds(len(m.items))
	return m.items[start:end]
}

func (m *listModel) ToggleSelectionAtCursor() {
	visible := m.CurrentPageItems()
	visible[m.cursor].Selected = !visible[m.cursor].Selected
}

func (m *listModel) CursorUp() {
	m.cursor--

	// We are at the very top of the list and cannot move up further
	if m.cursor < 0 && m.paginator.Page == 0 {
		m.cursor = 0
		return
	}

	// Move the cursor as normal
	if m.cursor >= 0 {
		return
	}

	// Go to the previous page
	m.paginator.PrevPage()
	m.cursor = m.paginator.ItemsOnPage(len(m.items)) - 1
}

func (m *listModel) CursorDown() {
	m.cursor++
	itemsOnPage := m.paginator.ItemsOnPage(len(m.items))

	// We are still within the page, move normally
	if m.cursor < itemsOnPage {
		return
	}

	// We are at the end of this page, but there are more pages. Move to the next page.
	if !m.paginator.OnLastPage() {
		m.paginator.NextPage()
		m.cursor = 0
		return
	}

	// We are at the end of the list and cannot move down further
	m.cursor = itemsOnPage - 1
}
