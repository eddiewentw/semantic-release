package main

import (
	"github.com/eddiewentw/semantic-release/internal/logger"
	"github.com/eddiewentw/semantic-release/pkg/file"
	"github.com/eddiewentw/semantic-release/pkg/flag"
	"github.com/eddiewentw/semantic-release/pkg/git"
	"github.com/eddiewentw/semantic-release/pkg/version"
)

func main() {
	args := flag.Parse()

	if args.IsFirstRelease == true {
		logger.DebugLog("first release", args.IsDebug)

		if err := file.WriteVersion(version.DEFAULT_VERSION); err != nil {
			logger.Error(err)
			return
		}

		if err := git.CommitRelease(version.DEFAULT_VERSION); err != nil {
			logger.Error(err)
		}

		return
	}

	currentTagVersion, err := git.GetLatestTagOnCurrentBranch()

	logger.DebugLog("current version is "+currentTagVersion, args.IsDebug)

	if err != nil {
		logger.Warning("no first release")
		return
	}

	commits, err := git.LogCommitsSince(currentTagVersion)

	logger.DebugLog("\n"+string(commits), args.IsDebug)

	if err != nil {
		logger.Error(err)
		return
	}

	nextVersion := version.Bump(currentTagVersion, commits)

	if args.IsDryRun == true {
		logger.Log("version: " + nextVersion + " (" + flag.DRY_RUN_FLAG + ")")
		return
	}

	if err = file.WriteVersion(nextVersion); err != nil {
		logger.Error(err)
		return
	}

	if err = git.CommitRelease(nextVersion); err != nil {
		logger.Error(err)
	}
}
