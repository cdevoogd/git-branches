package tui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

var ErrModelConversion = errors.New("failed to convert final model")

func PromptForSelections(choices []*Choice) ([]*Choice, error) {
	model := newSelectionModel(choices)
	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	final, ok := tm.(*selectionModel)
	if !ok || final == nil {
		return nil, ErrModelConversion
	}

	return final.Selections(), nil
}
