package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	greeting := strings.TrimSpace(remark)

	switch {
	case greeting == "":
		return "Fine. Be that way!"
	case any(greeting, unicode.IsUpper) && !any(greeting, unicode.IsLower):
		return "Whoa, chill out!"
	case greeting[len(greeting)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func any(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
