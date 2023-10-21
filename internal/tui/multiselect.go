package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	mapset "github.com/deckarep/golang-set/v2"
)

type selectionModel struct {
	cursor   int
	choices  []*Choice
	selected mapset.Set[int]
	keys     KeyMap
	help     help.Model
}

func newSelectionModel(choices []*Choice) *selectionModel {
	return &selectionModel{
		cursor:   0,
		choices:  choices,
		selected: mapset.NewSet[int](),
		keys:     DefaultKeyMap,
		help:     help.New(),
	}
}

func (m *selectionModel) Init() tea.Cmd {
	return nil
}

func (m *selectionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

		case key.Matches(msg, m.keys.Select):
			m.toggleSelection(m.cursor)

		case key.Matches(msg, m.keys.Confirm):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *selectionModel) View() string {
	view := &strings.Builder{}

	for i, choice := range m.choices {
		ctx := &choiceRenderContext{
			hovered:        m.cursor == i,
			selected:       m.isSelected(i),
			normalPrefix:   "[ ] ",
			hoveredPrefix:  "[•] ",
			selectedPrefix: "[x] ",
		}
		view.WriteString(choice.render(ctx))
		view.WriteString("\n")
	}

	view.WriteString("\n")
	view.WriteString(m.help.View(m.keys))
	return view.String()
}

func (m *selectionModel) Selections() []*Choice {
	iter := m.selected.Iterator()
	defer iter.Stop()

	var choices []*Choice
	for index := range iter.C {
		choices = append(choices, m.choices[index])
	}
	return choices
}

func (m *selectionModel) toggleSelection(index int) {
	if m.selected.Contains(index) {
		m.selected.Remove(index)
		return
	}
	m.selected.Add(index)
}

func (m *selectionModel) isSelected(index int) bool {
	return m.selected.Contains(index)
}
