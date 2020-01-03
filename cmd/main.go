package main

import (
	"github.com/eddiewentw/semantic-release/internal/logger"
	"github.com/eddiewentw/semantic-release/pkg/flag"
	"github.com/eddiewentw/semantic-release/pkg/git"
	"github.com/eddiewentw/semantic-release/pkg/version"
)

func main() {
	args := flag.Parse()

	if args.IsFirstRelease == true {
		if err := git.TagHead(version.DEFAULT_VERSION); err != nil {
			logger.Error(err)
		}

		return
	}

	currentTagVersion, err := git.GetLatestTagOnCurrentBranch()

	if err != nil {
		logger.Warning("no first release")
		return
	}

	commits, err := git.LogCommitsSince(currentTagVersion)

	if err != nil {
		logger.Error(err)
		return
	}

	nextVersion := version.Bump(currentTagVersion, commits)

	if args.IsDryRun == true {
		logger.Log("version: " + nextVersion + " (" + flag.DRY_RUN_FLAG + ")")
		return
	}

	if err = git.TagHead(nextVersion); err != nil {
		logger.Error(err)
	}
}
