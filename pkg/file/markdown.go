package file

func composeLink(text, url string) string {
	return "[" + text + "]" + "(" + url + ")"
}

func bold(text string) string {
	return "**" + text + "**"
}
