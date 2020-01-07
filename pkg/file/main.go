package file

import "io/ioutil"

const filename = ".semantic-version"
const Filepath = "./" + filename

func WriteVersion(version string) error {
	return ioutil.WriteFile(
		Filepath,
		[]byte(version+"\n"),
		0644,
	)
}
