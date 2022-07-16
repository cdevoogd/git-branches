package git

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type BranchType int

const (
	BranchTypeNormal BranchType = iota
	BranchTypeCurrent
	BranchTypeWorktree
)

type Branch struct {
	Type        BranchType
	Name        string
	Description string
	LastCommit  string
}

// NewBranch constructs a new Branch structure based on the given branch. The branch name must
// match the name of a branch that exists in the current Git repository. Using the name, the
// function will gather information about the branch and populate the structure.
func NewBranch(name string) (*Branch, error) {
	branchType, name := determineBranchType(name)
	desc := getBranchDescription(name)
	lastCommit := getLastCommit(name)

	return &Branch{
		Type:        branchType,
		Name:        name,
		Description: desc,
		LastCommit:  lastCommit,
	}, nil
}

// Branches returns a list of Git branches in the current repository.
func Branches() ([]*Branch, error) {
	cmd := exec.Command("git", "branch", "--list")
	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	output := strings.TrimSpace(string(stdout))

	var branches []*Branch
	for _, branch := range strings.Split(output, "\n") {
		b, err := NewBranch(branch)
		if err != nil {
			return nil, errors.Wrap(err, "error constructing branch")
		}
		branches = append(branches, b)
	}

	return branches, nil
}

const (
	worktreeBranchPrefix = "+"
	currentBranchPrefix  = "*"
)

// determineBranchType returns the type of the branch using the prefixes in its name as returned by
// the `git branch` command. The function also returns the name back to the caller with all
// prefixes and whitespace trimmed.
func determineBranchType(name string) (branchType BranchType, strippedName string) {
	switch {
	case strings.HasPrefix(name, currentBranchPrefix):
		branchType = BranchTypeCurrent
		name = strings.TrimPrefix(name, currentBranchPrefix)
	case strings.HasPrefix(name, worktreeBranchPrefix):
		branchType = BranchTypeWorktree
		name = strings.TrimPrefix(name, worktreeBranchPrefix)
	default:
		branchType = BranchTypeNormal
	}

	return branchType, strings.TrimSpace(name)
}

func getBranchDescription(name string) string {
	cmd := exec.Command("git", "config", "--get", "branch."+name+".description")
	stdout, _ := cmd.Output()

	// Git returns a non-zero exit code if the branch does not have a description, so I am relying
	// on the exit code instead of the returned error.
	if cmd.ProcessState.ExitCode() != 0 {
		return ""
	}

	return strings.TrimSpace(string(stdout))
}

func getLastCommit(branch string) string {
	cmd := exec.Command("git", "log", "-1", "--format=format:[%h] %s (%ah)", branch)
	stdout, _ := cmd.Output()

	// Git returns a non-zero exit code if the branch does not have a commit, so I am relying on
	// the exit code instead of the returned error.
	if cmd.ProcessState.ExitCode() != 0 {
		return ""
	}

	return strings.TrimSpace(string(stdout))
}
