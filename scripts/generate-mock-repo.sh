#!/usr/bin/env bash
set -euo pipefail

DEFAULT_REPO_DIR="/tmp/git-branches-mock"
DEFAULT_BRANCH_COUNT=30

REPO_DIR="$DEFAULT_REPO_DIR"
BRANCH_COUNT="$DEFAULT_BRANCH_COUNT"
WORKTREE_COUNT=3

print_usage() {
    echo "Generate a local git directory that can be used to manually test or debug git-branches"
    echo
    echo "USAGE:"
    echo "  $(basename "$0") [OPTIONS...]"
    echo
    echo "OPTIONS:"
    echo "  -h, --help      Print this help message"
    echo "  -d, --repo-dir  Override the directory used to create the repository (default: $DEFAULT_REPO_DIR)"
    echo "  -b, --branches  Set the number of branches to generate in the repo (default: $DEFAULT_BRANCH_COUNT)"
}

handle_missing_arg() {
    echo "Missing argument: $1"
    print_usage
    exit 1
}

parse_arguments() {
    while [[ $# -ne 0 ]] && [[ "$1" != "" ]]; do
        case $1 in
        -h | --help)
            print_usage
            exit
            ;;
        -d | --repo-dir)
            shift
            if [[ $# -eq 0 ]]; then handle_missing_arg "repo dir"; fi
            REPO_DIR="$1"
            ;;
        -b | --branches)
            shift
            if [[ $# -eq 0 ]]; then handle_missing_arg "branch count"; fi
            BRANCH_COUNT="$1"
            ;;
        *)
            echo "Unknown argument: $1"
            print_usage
            exit 1
            ;;
        esac
        shift
    done
}

log() {
    echo "[*]" "$@"
}

setup_repo_dir() {
    log "Setting up the mock repository directory: $REPO_DIR"
    if [[ -d "$REPO_DIR" ]]; then rm -rf "$REPO_DIR"; fi
    mkdir -p "$REPO_DIR"
}

init_git_repo() {
    git init --initial-branch main
    echo "# Mock Repository" > README.md
    echo "worktree-*" > .gitignore
    git add -A
    git commit -m "Initial commit"
}

create_branches() {
    log "Generating $BRANCH_COUNT branches"
    for i in $(seq 1 "$BRANCH_COUNT"); do
        git branch "branch-$i"
    done
}

create_worktrees() {
    log "Generating $WORKTREE_COUNT worktrees"
    for i in $(seq 1 "$WORKTREE_COUNT"); do
        git worktree add "worktree-$i"
    done
}

main() {
    parse_arguments "$@"
    setup_repo_dir
    cd "$REPO_DIR"
    init_git_repo
    create_branches
    create_worktrees
    log "Repo created: $REPO_DIR"
}

main "$@"
