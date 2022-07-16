package git

import (
	"errors"
	"os/exec"
	"regexp"
	"strings"
)

var notInRepoRegex = regexp.MustCompile("not a git repository")
var ErrNotInRepo = errors.New("the current working directory is not in a git repository")

// Branches returns a list of Git branches in the current repository.
func Branches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	stdout, err := cmd.Output()
	if err != nil {
		exitErr := &exec.ExitError{}
		ok := errors.As(err, &exitErr)
		if ok && notInRepoRegex.Match(exitErr.Stderr) {
			return nil, ErrNotInRepo
		}

		return nil, err
	}

	branches := strings.Split(string(stdout), "\n")
	return branches, nil
}
