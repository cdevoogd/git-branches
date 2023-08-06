package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cdevoogd/git-branches/internal/commands/delbranches"
	"github.com/cdevoogd/git-branches/internal/commands/listbranches"
	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
)

// Command line flags
var (
	deleteMode  bool
	versionMode bool
)

// These variables will be set during release builds by GoReleaser
// https://goreleaser.com/cookbooks/using-main.version/
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func printVersion() {
	fmt.Printf("notify %s\ncommit: %s\nbuild date: %s\n", version, commit, date)
}

func main() {
	flag.BoolVar(&deleteMode, "d", false, "Open a TUI to delete branches")
	flag.BoolVar(&versionMode, "version", false, "Print version information")
	flag.Parse()

	if versionMode {
		printVersion()
		os.Exit(0)
	}

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
