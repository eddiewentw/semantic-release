package git

import (
	"os/exec"
	"strings"
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

func TagHead(version string) error {
	return exec.Command("git", "tag", "-a", version, "-m", "chore(release): "+version).
		Run()
}
