package tui

import (
	"strings"

	"github.com/cdevoogd/git-branches/internal/tui/style"
	tea "github.com/charmbracelet/bubbletea"
)

type promptState int

const (
	promptStateNormal promptState = iota
	promptStateFinished
	promptStateError
)

type selectionModel interface {
	tea.Model
	Error() error
	Finished() bool
	Selections() []*Choice
}

type promptModel struct {
	model   selectionModel
	message string
}

func newPromptModel(message string, model selectionModel) *promptModel {
	return &promptModel{
		message: message,
		model:   model,
	}
}

func (p *promptModel) Init() tea.Cmd {
	return nil
}

func (p *promptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := p.model.Update(msg)
	p.model = model.(selectionModel)
	return p, cmd
}

func (p *promptModel) View() string {
	state := p.getModelState()

	s := strings.Builder{}

	switch state {
	case promptStateNormal:
		s.WriteString(style.NormalPrompt.Render("?"))
	case promptStateFinished:
		s.WriteString(style.SuccessPrompt.Render("✔"))
	case promptStateError:
		s.WriteString(style.ErrorPrompt.Render("✖"))
	}

	s.WriteString(" ")
	s.WriteString(p.message)
	s.WriteString(" ")

	if state == promptStateNormal {
		s.WriteString("\n")
		s.WriteString(p.model.View())
	} else {
		s.WriteString(style.NormalPrompt.Render("…"))
		s.WriteString(" ")
		s.WriteString(p.formatSelections())
		s.WriteString("\n")
	}

	return s.String()
}

func (p *promptModel) getModelState() promptState {
	if p.model.Error() != nil {
		return promptStateError
	}
	if p.model.Finished() {
		return promptStateFinished
	}
	return promptStateNormal
}

func (p *promptModel) formatSelections() string {
	selections := p.model.Selections()
	names := make([]string, len(selections))
	for i, choice := range selections {
		names[i] = choice.Name
	}
	return strings.Join(names, ", ")
}
