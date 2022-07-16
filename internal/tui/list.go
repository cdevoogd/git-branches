package tui

import (
	"fmt"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// item is a minimal implementation of the list's Item interface.
type listItem struct {
	title string
	desc  string
}

func (i *listItem) Title() string       { return i.title }
func (i *listItem) Description() string { return i.desc }
func (i *listItem) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func Start(branches []*git.Branch) error {
	items := make([]list.Item, len(branches))
	for i, b := range branches {
		items[i] = &listItem{
			title: b.Name,
			desc:  fmt.Sprintf("%s\n%s", b.Description, b.LastCommit),
		}
	}

	// Since the list item's description field contains the branch description and the last commit,
	// the delegate has a height of 3 to compensate (those two plus the branch name)
	delegate := list.NewDefaultDelegate()
	delegate.SetHeight(3)

	m := model{list: list.New(items, delegate, 0, 0)}
	m.list.Title = "Git Branches"

	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}
