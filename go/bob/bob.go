package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

func Hey(sentence string) string {
	sentence = strings.TrimSpace(sentence)
	switch {
	case len(sentence) == 0:
		return "Fine. Be that way!"
	case hasLetters(sentence) && isYelling(sentence):
		return "Whoa, chill out!"
	case strings.HasSuffix(sentence, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isYelling(s string) bool {
	for _, i := range s {
		if unicode.IsLower(i) {
			return false
		}
	}
	return true
}

func hasLetters(s string) bool {
	for _, i := range s {
		if unicode.IsLetter(i) {
			return true
		}
	}
	return false
}
