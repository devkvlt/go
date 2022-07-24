package rune

func IsHex(r rune) bool {
	return (47 < r && r < 58) || (96 < r && r < 103) || (64 < r && r < 71)
}
