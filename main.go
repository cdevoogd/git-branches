package main

import (
	"fmt"
	"os"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/tui"
)

func main() {
	branches, err := git.Branches()
	if err != nil {
		if git.ErrNotInRepository(err) {
			fmt.Println("The current directory is not part of a Git repository")
			os.Exit(1)
		}

		fmt.Println("Error querying branches:", err)
		os.Exit(1)
	}

	err = tui.Start(branches)
	if err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
