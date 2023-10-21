package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type singleSelectModel struct {
	cursor    int
	choices   []*Choice
	selection *Choice
	keys      SingleSelectKeyMap
	help      help.Model
	finished  bool
	err       error
}

func newSingleSelectModel(choices []*Choice) *singleSelectModel {
	return &singleSelectModel{
		cursor:  0,
		choices: choices,
		keys:    DefaultSingleSelectKeyMap,
		help:    help.New(),
	}
}

func (m *singleSelectModel) Init() tea.Cmd {
	return nil
}

func (m *singleSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Prev):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case key.Matches(msg, m.keys.Next):
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case key.Matches(msg, m.keys.Confirm):
			m.selection = m.choices[m.cursor]
			m.finished = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Quit):
			m.err = ErrQuit
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *singleSelectModel) View() string {
	view := &strings.Builder{}

	for i, choice := range m.choices {
		ctx := &choiceRenderContext{
			hovered: m.cursor == i,
		}
		view.WriteString(choice.render(ctx))
		view.WriteString("\n")
	}

	view.WriteString("\n")
	view.WriteString(m.help.View(m.keys))
	return view.String()
}

func (m *singleSelectModel) Selections() []*Choice {
	if m.selection == nil {
		return nil
	}
	return []*Choice{m.selection}
}

func (m *singleSelectModel) Finished() bool {
	return m.finished
}

func (m *singleSelectModel) Error() error {
	return m.err
}
