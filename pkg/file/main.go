package file

import (
	"io/ioutil"
	"strings"
)

const filename = ".semantic-version"
const Filepath = "./" + filename

func WriteVersion(version string) error {
	return ioutil.WriteFile(
		Filepath,
		[]byte(version+"\n"),
		0644,
	)
}

func ReadVersion() (string, error) {
	byteValue, err := ioutil.ReadFile(Filepath)

	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(byteValue), "\n"), nil
}
