package strings

import "github.com/devkvlt/go/runes"

func IsHex(s string) bool {
	for _, r := range s {
		if !runes.IsHex(r) {
			return false
		}
	}
	return true
}
