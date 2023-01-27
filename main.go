package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cdevoogd/git-branches/internal/git"
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

	for _, branch := range branches {
		bytes, err := json.MarshalIndent(branch, "", "    ")
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			os.Exit(1)
		}

		fmt.Println(string(bytes))
	}
}
