package tui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

var ErrModelConversion = errors.New("failed to convert final model")

func PromptForSelection(choices []*Choice) (*Choice, error) {
	model := newSelectModel(choices)
	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	final, ok := tm.(*singleSelectModel)
	if !ok || final == nil {
		return nil, ErrModelConversion
	}

	return final.Selection(), nil
}

func PromptForSelections(choices []*Choice) ([]*Choice, error) {
	model := newMultiSelectModel(choices)
	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	final, ok := tm.(*multiSelectModel)
	if !ok || final == nil {
		return nil, ErrModelConversion
	}

	return final.Selections(), nil
}
