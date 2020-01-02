package version

import (
	"regexp"
	"strconv"
	"strings"
)

type bumpLevel int

const (
	majorLevel bumpLevel = iota
	minorLevel bumpLevel = iota
	patchLevel bumpLevel = iota
)
const featureType = "feat"
const delimiter = "."

/*
	split git tag name into three parts, which are MAJOR, MINOR and PATCH
*/
func splitVersion(version string) (int, int, int) {
	/*
		there is a `v` prefix, remove it before splitting
	*/
	shards := strings.Split(version[1:], delimiter)

	major, err := strconv.Atoi(shards[0])

	if err != nil {
		panic(err)
	}

	minor, err := strconv.Atoi(shards[1])

	if err != nil {
		panic(err)
	}

	patch, err := strconv.Atoi(shards[2])

	if err != nil {
		panic(err)
	}

	return major, minor, patch
}

/*
	compose three parts to an entire version
*/
func composeVersion(major, minor, patch int) string {
	return "v" + strconv.Itoa(major) + delimiter + strconv.Itoa(minor) + delimiter + strconv.Itoa(patch)
}

const featureTypePrefix = featureType + "("

var gitMessageHeaderTypeRegex *regexp.Regexp

/*
	according commit messages, decide to how to bump versions
*/
func checkBumpLevel(commits []byte) bumpLevel {
	if gitMessageHeaderTypeRegex == nil {
		gitMessageHeaderTypeRegex = regexp.MustCompile(` ([a-z)(]+):`)
	}

	for _, matched := range gitMessageHeaderTypeRegex.FindAllSubmatch(commits, -1) {
		category := string(matched[1])
		if category == featureType || strings.HasPrefix(category, featureTypePrefix) {
			return minorLevel
		}
	}

	return patchLevel
}

func Bump(version string, commits []byte) string {
	major, minor, patch := splitVersion(version)
	level := checkBumpLevel(commits)

	if level == majorLevel {
		return composeVersion(major+1, 0, 0)
	}

	if level == minorLevel {
		return composeVersion(major, minor+1, 0)
	}

	return composeVersion(major, minor, patch+1)
}
