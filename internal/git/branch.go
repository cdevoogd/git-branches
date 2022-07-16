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

	lastCommit, err := getLastCommit(name)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting last commit for %s", name)
	}

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

	// If a branch doesn't have a description, Git will exit with an error. Because of that, I am
	// relying on the exit code alone instead of checking the returned error.
	stdout, _ := cmd.Output()
	if cmd.ProcessState.ExitCode() == 0 {
		return strings.TrimSpace(string(stdout))
	}

	return ""
}

func getLastCommit(branch string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--format=format:[%h] %s (%ah)", branch)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(stdout)), nil
}
