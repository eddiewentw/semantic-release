package file

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	changelogFilename = "CHANGELOG.md"
	ChangelogFilepath = "./" + changelogFilename
)
const fileHeader = `# Changelog

<!--- generated by semantic-release; DO NOT edit -->` + twoNewLines
const twoNewLines = "\n\n"

func generateChangedSection(commits []byte, regex *regexp.Regexp) string {
	text := ""

	for _, matched := range regex.FindAllSubmatch(commits, -1) {
		text = text + "- "

		if len(matched[2]) > 0 {
			text = text + "*" + string(matched[2]) + "*" + ": "
			text = strings.Replace(text, "(", "", 1)
			text = strings.Replace(text, ")", "", 1)
		}

		text = text + string(matched[3]) + " (" + string(matched[1]) + ")" + "\n"
	}

	if text == "" {
		return ""
	}

	return text + "\n"
}

const commitRegex = `(\w{7}) `
const messageRegex = `(\(.+\))?: (.+)`

var featMessageRegex = regexp.MustCompile(commitRegex + "feat" + messageRegex)

func generateFeatureSection(commits []byte) string {
	content := generateChangedSection(commits, featMessageRegex)

	if content == "" {
		return ""
	}

	return "#### New features" +
		twoNewLines +
		content
}

var fixMessageRegex = regexp.MustCompile(commitRegex + "fix" + messageRegex)

func generateFixSection(commits []byte) string {
	content := generateChangedSection(commits, fixMessageRegex)

	if content == "" {
		return ""
	}

	return "#### Fixed bugs" +
		twoNewLines +
		content
}

/*
	head of this change log section
*/
func generateVersionSection(version string) string {
	sectionMarkdown := "###"

	/*
		If new version ends with ".0", it bumped minor just now.
	*/
	if strings.HasSuffix(version, ".0") {
		sectionMarkdown = "##"
	}

	return sectionMarkdown + " " + version + twoNewLines
}

func readChangeLog() (string, error) {
	byteValue, err := ioutil.ReadFile(ChangelogFilepath)

	if err != nil {
		return "", err
	}

	return string(byteValue), nil
}

func WriteChangelog(version string, commits []byte) error {
	/*
		body: new section
	*/
	content := generateVersionSection(version)

	content = content + generateFeatureSection(commits)
	content = content + generateFixSection(commits)

	/*
		footer: old sections
	*/
	existedContent, err := readChangeLog()

	if existedContent == "" || err != nil {
		content = strings.TrimSuffix(content, "\n")
	} else {
		content = content + strings.TrimPrefix(existedContent, fileHeader)
	}

	/*
		header: file title
	*/
	content = fileHeader + content

	return ioutil.WriteFile(
		ChangelogFilepath,
		[]byte(content),
		0644,
	)
}
