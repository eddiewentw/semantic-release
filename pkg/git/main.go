package git

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/eddiewentw/semantic-release/internal/logger"
	"github.com/eddiewentw/semantic-release/pkg/constant"
)

/*
	returns a full git tag name, like `v1.0.4`
*/
func GetLatestTagOnCurrentBranch() (string, error) {
	out, err := exec.Command("git", "describe", "--abbrev=0").
		Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

func LogCommitsSince(tag string) ([]byte, error) {
	return exec.Command("git", "log", tag+".."+"HEAD", "--oneline").
		Output()
}

func CommitRelease(version string) error {
	err := exec.Command("git", "add", constant.VersionFilepath, constant.ChangelogFilepath).
		Run()

	if err != nil {
		return err
	}

	err = exec.Command("git", "commit", "-m", "chore(release): "+version).
		Run()

	if err != nil {
		return err
	}

	logger.Log("version: " + version)

	return exec.Command("git", "tag", "-a", version, "-m", "chore(release): "+version).
		Run()
}

var protocolRegex = regexp.MustCompile(".*@")

func GetRepoUrl() string {
	out, err := exec.Command("git", "remote", "get-url", "origin").
		Output()

	if err != nil {
		return ""
	}

	url := strings.Replace(string(out), ".git", "", 1)
	url = strings.Replace(url, ":", "/", 1)
	url = protocolRegex.ReplaceAllString(url, "https://")
	url = strings.TrimSuffix(url, "\n")

	return url
}
