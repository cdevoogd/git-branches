package git

import (
	"errors"
	"os/exec"
	"regexp"
)

var notInRepoRegex = regexp.MustCompile("not a git repository")

// ErrNotInRepository returns true if the given error occurred because the current working 
// directory is not part of a Git repository.
func ErrNotInRepository(err error) bool {
	// The exec package used to run Git commands should return errors that are ExitError types
	exitErr := &exec.ExitError{}
	ok := errors.As(err, &exitErr)
	if !ok {
		return false
	}
	return notInRepoRegex.Match(exitErr.Stderr)
}