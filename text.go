package loopfmt

func Bold(text string) string {
	return "**" + text + "**"
}

func Italic(text string) string {
	return "*" + text + "*"
}

func Strikethrough(text string) string {
	return "~~" + text + "~~"
}

func Header(text string) string {
	return "###" + text
}

func Link(items ...string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return "[" + items[0] + "]"
	}
	return "[" + items[0] + "](" + items[1] + ")"
}

func CodeBlock(code string) string {
	return "```" + code + "```"
}

func Quote(block string) string {
	return "> " + block
}
