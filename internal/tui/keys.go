package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Prev    key.Binding
	Next    key.Binding
	Select  key.Binding
	Confirm key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Select, k.Confirm, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next, k.Select},  // first column
		{k.Help, k.Confirm, k.Quit}, // second column
	}
}

var DefaultKeyMap = KeyMap{
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
