package main

import (
	"github.com/eddiewentw/semantic-release/pkg/git"
	"github.com/eddiewentw/semantic-release/pkg/version"
)

func main() {
	currentTagVersion, err := git.GetLatestTagOnCurrentBranch()

	if err != nil {
		panic(err)
	}

	commits, err := git.LogCommitsSince(currentTagVersion)

	if err != nil {
		panic(err)
	}

	nextVersion := version.Bump(currentTagVersion, commits)

	err = git.TagHead(nextVersion)

	if err != nil {
		panic(err)
	}
}
