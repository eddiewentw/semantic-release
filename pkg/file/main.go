package file

import (
	"io/ioutil"
	"strings"

	"github.com/eddiewentw/semantic-release/pkg/constant"
)

func WriteVersion(version string) error {
	return ioutil.WriteFile(
		constant.VersionFilepath,
		[]byte(version+"\n"),
		0644,
	)
}

func ReadVersion() (string, error) {
	byteValue, err := ioutil.ReadFile(constant.VersionFilepath)

	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(byteValue), "\n"), nil
}
