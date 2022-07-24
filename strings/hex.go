package strings

import "github.com/devkvlt/go/runes"

func IsHex[T ~string](s T) bool {
	for _, r := range s {
		if !runes.IsHex(r) {
			return false
		}
	}
	return true
}
