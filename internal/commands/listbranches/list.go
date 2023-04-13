package listbranches

import (
	"fmt"

	"github.com/cdevoogd/git-branches/internal/git"
	"github.com/cdevoogd/git-branches/internal/log"
	"github.com/fatih/color"
)

type branchStyle struct {
	*color.Color
	prefix string
}

var (
	descStyle  = color.New(color.FgWhite)
	nameStyles = map[git.BranchType]branchStyle{
		git.BranchTypeNormal:   {prefix: "  ", Color: color.New(color.Bold, color.FgWhite)},
		git.BranchTypeCurrent:  {prefix: "* ", Color: color.New(color.Bold, color.FgGreen)},
		git.BranchTypeWorktree: {prefix: "+ ", Color: color.New(color.Bold, color.FgCyan)},
	}
)

// Run will print out information (name, type, desc) about the given branches to stdout.
func Run(branches []*git.Branch) int {
	for _, branch := range branches {
		err := printBranch(branch)
		if err != nil {
			log.Error(err)
			return 1
		}
	}

	return 0
}

func printBranch(branch *git.Branch) error {
	nameStyle, ok := nameStyles[branch.Type]
	if !ok {
		return fmt.Errorf("no style is available for branch type %q", branch.Type)
	}

	name := nameStyle.Sprint(nameStyle.prefix, branch.Name)
	desc := branch.Description
	if desc != "" {
		fmt.Printf("%s (%s)\n", name, descStyle.Sprint(desc))
		return nil
	}

	fmt.Println(name)
	return nil
}
