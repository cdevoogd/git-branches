package delete

import (
	"strings"

	"github.com/cqroot/prompt"
)

// displayBranches returns a styled string that can be used to display the prompt to the user. It
// is meant to fulfil the prompt.Theme interface. This acts similar to the default theme but with
// colors that match the rest of the branch choice prompt.
func themePrompt(msg string, state prompt.State, model string) string {
	s := strings.Builder{}

	switch state {
	case prompt.StateNormal:
		s.WriteString(defaultPromptStyle.Render("?"))
	case prompt.StateFinish:
		s.WriteString(finishPromptPrefixStyle.Render("✔"))
	case prompt.StateError:
		s.WriteString(defaultPromptStyle.Render("✖"))
	}

	s.WriteString(" ")
	s.WriteString(msg)
	s.WriteString(" ")

	if state == prompt.StateNormal {
		s.WriteString(model)
	} else {
		s.WriteString(defaultPromptStyle.Render("…"))
		s.WriteString(" ")
		s.WriteString(model)
		s.WriteString("\n")
	}

	return s.String()
}
