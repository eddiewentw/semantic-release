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

	/*
		figure out current version
	*/
	currentVersion, err := file.ReadVersion()

	if err != nil {
		currentVersion, err = git.GetLatestTagOnCurrentBranch()

		if err != nil {
			logger.Warning("no first release")
			return
		}
	}

	logger.DebugLog("current version is "+currentVersion, args.IsDebug)
	commits, err := git.LogCommitsSince(currentVersion)

	logger.DebugLog("\n"+string(commits), args.IsDebug)

	if err != nil {
		logger.Error(err)
		return
	}

	nextVersion := version.Bump(currentVersion, commits)

	if err = file.WriteChangelog(nextVersion, commits); err != nil {
		logger.Error(err)
		return
	}

	if err = file.WriteVersion(nextVersion); err != nil {
		logger.Error(err)
		return
	}

	if args.IsDryRun == true {
		logger.Log("version: " + nextVersion + " (" + flag.DRY_RUN_FLAG + ")")
		return
	}

	if err = git.CommitRelease(nextVersion); err != nil {
		logger.Error(err)
	}
}
