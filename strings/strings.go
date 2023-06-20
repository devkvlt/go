package strings

import "github.com/devkvlt/go/runes"

// IsHex reports whether a string is made of only hexadecimal characters.
func IsHex[T ~string](s T) bool {
	for _, r := range s {
		if !runes.IsHex(r) {
			return false
		}
	}
	return true
}
