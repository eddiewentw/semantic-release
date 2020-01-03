package main

import (
	"fmt"

	"github.com/eddiewentw/semantic-release/pkg/flag"
	"github.com/eddiewentw/semantic-release/pkg/git"
	"github.com/eddiewentw/semantic-release/pkg/version"
)

func main() {
	args := flag.Parse()

	if args.IsFirstRelease == true {
		if err := git.TagHead(version.DEFAULT_VERSION); err != nil {
			panic(err)
		}

		return
	}

	currentTagVersion, err := git.GetLatestTagOnCurrentBranch()

	if err != nil {
		panic(err)
	}

	commits, err := git.LogCommitsSince(currentTagVersion)

	if err != nil {
		panic(err)
	}

	nextVersion := version.Bump(currentTagVersion, commits)

	fmt.Print("Release:", nextVersion)

	if args.IsDryRun == true {
		fmt.Println(" (" + flag.DRY_RUN_FLAG + ")")
		return
	}

	fmt.Print("\n")

	if err = git.TagHead(nextVersion); err != nil {
		panic(err)
	}
}
