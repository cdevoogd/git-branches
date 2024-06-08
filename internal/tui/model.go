package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	whiteStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	greenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6EDD8C"))
	cyanStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#7DF9FE"))

	listNormalBranchStyle   = whiteStyle.SetString(" ")
	listCurrentBranchStyle  = greenStyle.SetString("*")
	listWorktreeBranchStyle = cyanStyle.SetString("+")

	multiSelectUnselectedStyle         = lipgloss.NewStyle().SetString("[ ]")
	multiSelectSelectedStyle           = lipgloss.NewStyle().SetString("[x]")
	multiSelectHoveredStyle            = greenStyle.SetString("[•]")
	multiSelectHoveredAndSelectedStyle = greenStyle.SetString().SetString("[x]")

	selectHoveredStyle = greenStyle.SetString(">")
	selectDefaultStyle = lipgloss.NewStyle().SetString(" ")

	descriptionStyle           = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"}).SetString(" ")
	activePaginationDotStyle   = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	inactivePaginationDotStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
)

type MultiSelectModel struct {
	items    []*Item
	finished bool
	err      error

	keys MultiSelectKeyMap
	list listModel
	help help.Model
}

func NewMultiSelectModel(items []*Item) *MultiSelectModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = activePaginationDotStyle
	p.InactiveDot = inactivePaginationDotStyle
	p.PerPage = 10
	p.SetTotalPages(len(items))

	return &MultiSelectModel{
		items: items,
		keys:  DefaultMultiSelectKeyMap,
		list:  *newListModel(items, renderMultiSelectItem),
		help:  help.New(),
	}
}

func (m *MultiSelectModel) Init() tea.Cmd {
	return nil
}

func (m *MultiSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Prev):
			m.list.CursorUp()

		case key.Matches(msg, m.keys.Next):
			m.list.CursorDown()

		case key.Matches(msg, m.keys.Select):
			m.list.ToggleSelectionAtCursor()

		case key.Matches(msg, m.keys.Confirm):
			m.finished = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Quit):
			m.finished = true
			m.err = ErrQuit
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *MultiSelectModel) View() string {
	if m.finished {
		return ""
	}

	view := &strings.Builder{}
	view.WriteString(m.list.View())
	view.WriteRune('\n')
	view.WriteString(m.help.View(m.keys))
	return view.String()
}

func (m *MultiSelectModel) SelectedItems() []*Item {
	return m.list.SelectedItems()
}

func (m *MultiSelectModel) Error() error {
	return m.err
}

type SelectModel struct {
	items    []*Item
	finished bool
	err      error

	keys SingleSelectKeyMap
	list listModel
	help help.Model
}

func NewSingleSelectModel(items []*Item) *SelectModel {
	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = activePaginationDotStyle
	p.InactiveDot = inactivePaginationDotStyle
	p.PerPage = 10
	p.SetTotalPages(len(items))

	return &SelectModel{
		items: items,
		keys:  DefaultSingleSelectKeyMap,
		list:  *newListModel(items, renderSingleSelectItem),
		help:  help.New(),
	}
}

func (m *SelectModel) Init() tea.Cmd {
	return nil
}

func (m *SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Prev):
			m.list.CursorUp()

		case key.Matches(msg, m.keys.Next):
			m.list.CursorDown()

		case key.Matches(msg, m.keys.Confirm):
			m.list.ToggleSelectionAtCursor()
			m.finished = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Quit):
			m.finished = true
			m.err = ErrQuit
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *SelectModel) View() string {
	if m.finished {
		return ""
	}

	view := &strings.Builder{}
	view.WriteString(m.list.View())
	view.WriteRune('\n')
	view.WriteString(m.help.View(m.keys))
	return view.String()
}

func (m *SelectModel) SelectedItem() *Item {
	items := m.list.SelectedItems()
	if len(items) == 0 {
		return nil
	}
	return items[0]
}

func (m *SelectModel) Error() error {
	return m.err
}

func renderMultiSelectItem(item *Item, hovered bool) string {
	out := &strings.Builder{}
	switch {
	case hovered && item.Selected:
		out.WriteString(multiSelectHoveredAndSelectedStyle.Render(item.Name))
	case hovered:
		out.WriteString(multiSelectHoveredStyle.Render(item.Name))
	case item.Selected:
		out.WriteString(multiSelectSelectedStyle.Render(item.Name))
	default:
		out.WriteString(multiSelectUnselectedStyle.Render(item.Name))
	}

	if item.Note != "" {
		out.WriteString(descriptionStyle.Render(item.Note))
	}

	return out.String()
}

func renderSingleSelectItem(item *Item, hovered bool) string {
	out := &strings.Builder{}
	switch {
	case hovered:
		out.WriteString(selectHoveredStyle.Render(item.Name))
	default:
		out.WriteString(selectDefaultStyle.Render(item.Name))
	}

	if item.Note != "" {
		out.WriteString(descriptionStyle.Render(item.Note))
	}

	return out.String()
}
