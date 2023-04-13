package log

import (
	"fmt"
	"os"

	"github.com/cdevoogd/git-branches/internal/color"
	"github.com/charmbracelet/lipgloss"
)

var errorPrefix = lipgloss.NewStyle().Foreground(color.Red).Render("Error:")

// Error prints out an error-level log to stderr.
func Error(v ...any) {
	fmt.Fprintln(os.Stderr, errorPrefix, fmt.Sprint(v...))
}

// Errorf prints out an error-level log to strderr. Formatting arguments are handled as they would
// with fmt.Printf.
func Errorf(format string, v ...any) {
	fmt.Fprintln(os.Stderr, errorPrefix, fmt.Sprintf(format, v...))
}

// Fatal is equivalent to Error() followed by a call to os.Exit(1).
func Fatal(v ...any) {
	Error(v...)
	os.Exit(1)
}

// Fatalf is equivalent to Errorf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...any) {
	Errorf(format, v...)
	os.Exit(1)
}
