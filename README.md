# Git Branches
This is a simple tool to replace the usual `git branch` printout provided by Git.

When developing on branches that use ticket numbers, I find that it can sometimes be difficult to remember what exactly the branch was for. To help with this, I would add descriptions to the branches using `git branch --edit-description`, but Git does not provide an easy way for you to quickly look at all of the branches along with their descriptions. Instead, to get the description of a branch, you have to use a long command: `git config --get branch.<branch_name>.description`.

This tool provides an alternative to `git branch` that also displays branch descriptions (if available) and recent commits. I personally find that this makes it a lot easier to quickly scan over my local brances, so hopefully it can help others too.

## Installation
Currently, the best way to install this is directly with Go:
```shell
go install github.com/cdevoogd/git-branches@latest
```

## Notes
### Setting Descriptions

Adding descriptions to your branches when using this tool can make it much easier to quickly determine what you were working on on a branch. I recommend [setting an alias in your `.gitconfig`](https://github.com/cdevoogd/.dotfiles/blob/9f2d907c4afd5e4b2cbe8b4b0ea52d0dae5f9ddc/git/main.gitconfig#L7) to assist in quickly setting descriptions.

### Git Subcommand

If you leave the name of this program as-is when installing, Git should automatically allow it to be used as a custom subcommand. If the program is named `git-branches` and is available in your PATH, you should be able to call it using `git branches`.
