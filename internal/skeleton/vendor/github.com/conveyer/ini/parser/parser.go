// Package parser provides functions necessary for parsing
// INI configuration format.
package parser

import (
	"bufio"
	"fmt"
)

const (
	commentBeg  = '#'
	kvSeparator = '='
	sectionBeg  = '['
	sectionEnd  = ']'
	doubleQuote = '"'
)

// Section represents a section of INI file.
// It contains its name and keys along with values.
type Section struct {
	Name         []byte
	Keys, Values [][]byte
}

// context represents an instance of a single parser.
type context struct {
	sections []Section
	currLine int
}

// Parse gets some INI configuration as bufio.Scanner, transforms it
// into a Go object and returns. The result is a 1:1 representation,
// except comments are omitted.
// No assumptions are made about what to do with repeating keys or sections
// and other stuff like that intentionally, so this can be
// handled on a higher layer depending on requirements.
// If the requested configuration cannot be parsed
// a non-nil error will be returned as a second argument.
func Parse(s *bufio.Scanner) ([]Section, error) {
	// Handle the input line-by-line till the end
	// is reached.
	c := newContext()
	for s.Scan() {
		c.currLine++
		err := c.parseLine(s.Bytes())
		if err != nil {
			return nil, fmt.Errorf("ini syntax error on line %d: %s", c.currLine, err)
		}
	}

	// Make sure the input has been scanned correctly.
	if err := s.Err(); err != nil {
		return nil, err
	}

	// If no errors are returned so far, the input configuration
	// has been parsed successfully. Return the result.
	return c.sections, nil
}

// newContext allocates and returns a new context
// with the default section.
func newContext() *context {
	return &context{
		sections: []Section{{Name: []byte("")}}, // Default section's name is "".
	}
}

// add appends a new key-value pair to the section.
func (s *Section) add(k, v []byte) {
	s.Keys = append(s.Keys, k)
	s.Values = append(s.Values, v)
}
