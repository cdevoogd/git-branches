package main

import (
	"github.com/cdevoogd/git-branches/internal/commands/list"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
)

func main() {
	branches, err := git.Branches()
	if err != nil {
		if git.ErrNotInRepository(err) {
			log.Fatal("The current directory is not part of a Git repository")
		}

		log.Fatal("An error occured when querying for branches:", err)
	}

	err = list.PrintBranches(branches)
	if err != nil {
		log.Fatal("Error listing branches:", err)
	}
}
