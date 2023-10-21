package tui

import "github.com/charmbracelet/bubbles/key"

type MultiSelectKeyMap struct {
	Prev    key.Binding
	Next    key.Binding
	Select  key.Binding
	Confirm key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k MultiSelectKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Select, k.Confirm, k.Quit}
}

func (k MultiSelectKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next, k.Select},  // first column
		{k.Help, k.Confirm, k.Quit}, // second column
	}
}

var DefaultMultiSelectKeyMap = MultiSelectKeyMap{
	Prev: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Next: key.NewBinding(
		key.WithKeys("down", "j", "tab"),
		key.WithHelp("↓/j/tab", "move down"),
	),
	Select: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("space", "select"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type SingleSelectKeyMap struct {
	Prev    key.Binding
	Next    key.Binding
	Confirm key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k SingleSelectKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Confirm, k.Quit}
}

func (k SingleSelectKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next, k.Confirm}, // first column
		{k.Help, k.Quit},            // second column
	}
}

var DefaultSingleSelectKeyMap = SingleSelectKeyMap{
	Prev: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Next: key.NewBinding(
		key.WithKeys("down", "j", "tab"),
		key.WithHelp("↓/j/tab", "move down"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "confirm"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
