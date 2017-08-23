package pangram

import (
	"strings"
)

const testVersion = 2

func IsPangram(t string) bool {

	upper := strings.ToUpper(t)
	for l := 65; l <= 90; l++ {
		if !(strings.ContainsRune(upper, rune(l))) {
			return false
		}
	}
	return true
}
