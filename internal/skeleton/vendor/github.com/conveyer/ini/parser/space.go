package parser

import (
	"unicode"
)

// trimSpaceLeft returns a value without leading spaces.
// The length of the result is returned as a second argument.
func trimSpaceLeft(v []byte) ([]byte, int) {
	// If the value is empty, return it as is.
	l := len(v)
	if l == 0 {
		return v, l
	}

	// Ignore the leading spaces of the value.
	if unicode.IsSpace(rune(v[0])) {
		return trimSpaceLeft(v[1:])
	}

	// Return both the value and its length.
	return v, l
}
