package file

import (
	"io/ioutil"
	"strings"
)

const (
	versionFilename = ".semantic-version"
	VersionFilepath = "./" + versionFilename
)

func WriteVersion(version string) error {
	return ioutil.WriteFile(
		VersionFilepath,
		[]byte(version+"\n"),
		0644,
	)
}

func ReadVersion() (string, error) {
	byteValue, err := ioutil.ReadFile(VersionFilepath)

	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(byteValue), "\n"), nil
}
