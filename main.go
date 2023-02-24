package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
)

type branchStyle struct {
	*color.Color
	prefix string
}

var (
	descStyle   = color.New(color.FgWhite)
	commitStyle = color.New(color.Italic, color.FgWhite)
	nameStyles  = map[git.BranchType]branchStyle{
		git.BranchTypeNormal:   {prefix: "  ", Color: color.New(color.Bold, color.FgWhite)},
		git.BranchTypeCurrent:  {prefix: "* ", Color: color.New(color.Bold, color.FgGreen)},
		git.BranchTypeWorktree: {prefix: "+ ", Color: color.New(color.Bold, color.FgCyan)},
	}
)

func main() {
	branches, err := git.Branches()
	if err != nil {
		if git.ErrNotInRepository(err) {
			log.Fatal("The current directory is not part of a Git repository")
		}

		log.Fatal("An error occured when querying for branches:", err)
	}

	lastIndex := len(branches) - 1
	for i, branch := range branches {
		nameStyle, ok := nameStyles[branch.Type]
		if !ok {
			log.Fatalf("No style is available for branch type %q", branch.Type)
		}

		nameStyle.Printf("%s%s", nameStyle.prefix, branch.Name)
		if branch.Description != "" {
			descStyle.Printf(" (%s)", branch.Description)
		}
		commitStyle.Printf("\n    %s\n", branch.LastCommit)

		// Add some spacing between branches
		if i != lastIndex {
			fmt.Println()
		}
	}
}
