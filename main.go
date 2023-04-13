package main

import (
	"os"

	"github.com/cdevoogd/git-branches/internal/commands/delbranches"
	"github.com/cdevoogd/git-branches/internal/commands/list"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
)

const deleteCommand = "delete"

func shouldRunDelete() bool {
	return len(os.Args) == 2 && os.Args[1] == deleteCommand
}

func main() {
	branches, err := git.Branches()
	if err != nil {
		if git.ErrNotInRepository(err) {
			log.Fatal("The current directory is not part of a Git repository")
		}

		log.Fatal("An error occurred when querying for branches: ", err)
	}

	if shouldRunDelete() {
		os.Exit(delbranches.Run(branches))
	}

	os.Exit(list.PrintBranches(branches))
}
