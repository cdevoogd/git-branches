package main

import (
	"flag"
	"os"

	"github.com/cdevoogd/git-branches/internal/commands/delbranches"
	"github.com/cdevoogd/git-branches/internal/commands/listbranches"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
)

var deleteMode bool

func main() {
	flag.BoolVar(&deleteMode, "d", false, "Open a TUI to delete branches")
	flag.Parse()

	branches, err := git.Branches()
	if err != nil {
		if git.ErrNotInRepository(err) {
			log.Fatal("The current directory is not part of a Git repository")
		}

		log.Fatal("An error occurred when querying for branches: ", err)
	}

	if deleteMode {
		os.Exit(delbranches.Run(branches))
	}

	os.Exit(listbranches.Run(branches))
}
