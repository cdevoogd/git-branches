package main

import (
	"fmt"
	"os"

	"github.com/cdevoogd/git-branches/internal/git"
)

func main() {
	branches, err := git.Branches()
	if err != nil {
		fmt.Println("Error querying branches:", err)
		os.Exit(1)
	}

	fmt.Println(branches)
}
