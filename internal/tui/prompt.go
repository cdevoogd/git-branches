package tui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	ErrModelConversion = errors.New("failed to convert final model")
	ErrQuit            = errors.New("user exited prompt")
)

func PromptForSelection(message string, choices []*Choice) (*Choice, error) {
	singleSelect := newSingleSelectModel(choices)
	prompt := newPromptModel(message, singleSelect)

	tm, err := tea.NewProgram(prompt).Run()
	if err != nil {
		return nil, err
	}

	finalPrompt, ok := tm.(*promptModel)
	if !ok || finalPrompt == nil {
		return nil, ErrModelConversion
	}

	err = finalPrompt.model.Error()
	if err != nil {
		return nil, err
	}

	selections := finalPrompt.model.Selections()
	if len(selections) == 0 {
		return nil, nil
	}

	return selections[0], nil
}

func PromptForSelections(message string, choices []*Choice) ([]*Choice, error) {
	multiSelect := newMultiSelectModel(choices)
	prompt := newPromptModel(message, multiSelect)

	tm, err := tea.NewProgram(prompt).Run()
	if err != nil {
		return nil, err
	}

	finalPrompt, ok := tm.(*promptModel)
	if !ok || finalPrompt == nil {
		return nil, ErrModelConversion
	}

	err = finalPrompt.model.Error()
	if err != nil {
		return nil, err
	}

	return finalPrompt.model.Selections(), nil
}
