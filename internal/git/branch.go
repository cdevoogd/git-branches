package git

import (
	"os/exec"
	"strings"
)

type BranchType int

const (
	BranchTypeNormal BranchType = iota
	BranchTypeCurrent
	BranchTypeWorktree
)

type Branch struct {
	Type BranchType
	Name string
}

// Branches returns a list of Git branches in the current repository.
func Branches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(string(stdout), "\n")
	return branches, nil
}
