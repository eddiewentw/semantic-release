package git

import (
	"os"
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
	err := exec.Command("git", "add", constant.VersionFilepath).
		Run()

	if err != nil {
		return err
	}

	/*
		check if changelog file exists

		There is no changelog file when the first release.
	*/
	_, err = os.Stat(constant.ChangelogFilepath)

	if os.IsNotExist(err) == false {
		err := exec.Command("git", "add", constant.ChangelogFilepath).
			Run()

		if err != nil {
			return err
		}
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

const httpsProtocol = "https://"

func GetRepoURL() string {
	out, err := exec.Command("git", "remote", "get-url", "origin").
		Output()

	if err != nil {
		return ""
	}

	/*
		HTTPS protocol
	*/
	if strings.HasPrefix(string(out), httpsProtocol) == true {
		return strings.TrimSuffix(string(out), "\n")
	}

	/*
		SSH protocol
	*/
	url := strings.Replace(string(out), ".git", "", 1)
	url = strings.Replace(url, ":", "/", 1)
	url = protocolRegex.ReplaceAllString(url, httpsProtocol)
	url = strings.TrimSuffix(url, "\n")

	return url
}
