package parser

import (
	"fmt"
	"unicode"
)

// parseKV gets a fragment of INI configuration and extracts
// key and value out of it.
// Samples of correct input are:
//	key1 = value1
//	key2 = "   value2   "#Spaces around the value2 will be preserved.
//	key3 = \"Something here\"  # Double quotes will be preserved.
//	ключ =  \t какое-то значение # Leading and trailing spaces will be removed.
//	key4[] = "whatever"
//	"key5"=value5
//	key# = value
// No leading space is expected as it has already been cleaned.
func (c *context) parseKV(kv []byte) (k []byte, v []byte, err error) {
	// Looking for the end of the key.
	l := len(kv)
	endInd := l // By default, the end of the line is the end of the key.
	for i := range kv {
		switch currC := kv[i]; true {
		case unicode.IsSpace(rune(currC)):
			// If current character is a space and we haven't found
			// the end of key, assume that it is a trailing space.
			if endInd == l {
				endInd = i
			}
		case currC == kvSeparator:
			// If there were trailing spaces, use the last position before them,
			// otherwise, use the current position as the end of the key.
			if endInd == l {
				endInd = i
			}

			// Key-value separator has been found. That means
			// that the rest of the fragment is the value.
			v, err := c.parseValue(kv[i+1:])
			if err != nil {
				return nil, nil, err
			}
			return kv[:endInd], v, nil
		default:
			// If there were spaces before, they were not trailing.
			// So, restore the default position of the last element.
			endInd = l
		}
	}
	return nil, nil, fmt.Errorf(
		`"%c" separator is missing after the key "%s"`, kvSeparator, kv[:endInd],
	)
}

// parseValue gets a value fragment and parses it.
// Samples of the correct input include:
//	\t
//	value1
//	# Some comment
//	value=1
//	"  value  1  "
//	Hello, "world"
//	\"Something\"
func (c *context) parseValue(v []byte) ([]byte, error) {
	// Clean the trailing spaces.
	v, l := trimSpaceLeft(v)
	if l == 0 {
		return v, nil
	}

	// Find the beginning and the end of the value.
	startsWithQuote := v[0] == doubleQuote
	quoted := startsWithQuote
	begInd, endInd := 0, l
loop:
	for i := range v {
		switch currC := v[i]; true {
		case currC == commentBeg && !quoted:
			// Omit the comment.
			if endInd == l {
				endInd = i
			}
			break loop
		case unicode.IsSpace(rune(currC)) && !quoted:
			// If we haven't found the end of the value yet,
			// assume that the current space is trailing.
			if endInd == l {
				endInd = i
			}
			continue
		case currC == doubleQuote && startsWithQuote:
			// Ignore the first double quote character.
			if i == 0 {
				continue
			}

			// Disable the "quoted" mode.
			if quoted {
				quoted = false
				continue
			}

			// The value starts with a quote, but the "quoted"
			// mode is not active. That means that the quotes
			// have already been closed and now there is an attempt
			// to open them again.
			return nil, fmt.Errorf("string literal has already been terminated near `%s`", v[i:])
		case startsWithQuote && !quoted:
			// The value was started with a quote,
			// but after it is closed, some other characters
			// we don't know how to hadle are placed.
			goto unterminatedLiteral
		}

		// Restore the position of the last element.
		endInd = l
	}

	// Double quote characters should not be part
	// of the value.
	if startsWithQuote {
		begInd++
		endInd--
	}

	// If string literal has been closed correctly,
	// return the result value.
	if !quoted {
		return v[begInd:endInd], nil
	}

	// Otherwise, return an error informing about unterminated string literal.
unterminatedLiteral:
	return nil, fmt.Errorf("string literal of `%s` not terminated", v)
}
