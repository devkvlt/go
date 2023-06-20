package runes

// IsHex reports whether a rune is a hexadecimal.
func IsHex(r rune) bool {
	return ('0' <= r && r <= '9') || ('a' <= r && r <= 'f') || ('A' <= r && r <= 'F')
}
