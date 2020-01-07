package file

import "io/ioutil"

const filename = ".semantic-version"
const filepath = "./" + filename

func WriteVersion(version string) error {
	return ioutil.WriteFile(
		filepath,
		[]byte(version+"\n"),
		0644,
	)
}
