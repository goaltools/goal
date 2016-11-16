package parser

import (
	"errors"
	"fmt"
	"unicode"
)

// parseSection gets a section fragment, parses and returns it.
// Samples of correct inputs (the openning "[" part has already been processed):
//	sampleSection]
//	sampleSection]# Some comment
//	sampleSection]    # Some comment
//	sample[Section]]
//	[образец][][]Секции]
//	  #Какая-то = "секция"  ]
func (c *context) parseSection(section []byte) ([]byte, error) {
	// Make sure the section fragment is not empty.
	l := len(section)
	if l == 0 {
		return nil, fmt.Errorf(`incorrect section declaration, "%c" is missing`, sectionEnd)
	}

	// Ignore leading spaces of the section name.
	if unicode.IsSpace(rune(section[0])) {
		return c.parseSection(section[1:])
	}

	// Prepare for parsing of the actual section name.
	unclosedBr := 1 // unclosedBr stores a number of times section brackets have been opened.
	endInd := l     // endInd stores an index of the last element of actual section name.

	// Iterate over the section name's characters and parse them appropriately.
loop:
	for i := range section {
		switch currC := section[i]; true {
		case unicode.IsSpace(rune(currC)):
			// If we still haven't found the index of the actual section name's
			// last element and the current character is a space, let's assume
			// this space is trailing and thus considered the ending element for now.
			if endInd == l {
				endInd = i
			}
			continue
		case currC == sectionBeg:
			// Increment the number of unclosed section brackets.
			unclosedBr++
		case currC == sectionEnd:
			// Decrement the number of unclosed section brackets.
			unclosedBr--

			// If this is not the last bracket, do nothing special (break from the switch).
			if unclosedBr != 0 {
				break
			}

			// If all of the brackets have been closed but the last
			// element hasn't been set yet, do it now.
			if endInd == l {
				endInd = i
			}

			// Do not proceed with the current iteration so the ending index
			// is not overridden.
			continue
		case unclosedBr == 0 && currC == commentBeg:
			// If all of the section brackets are closed and the current
			// character indicates the beginning of a comment, ignore the rest.
			break loop
		case unclosedBr == 0:
			// All of the brackets are closed, but there are still some characters
			// we don't know how to handle. That means the input is not correct.
			return nil, fmt.Errorf(`error near "%s", section name cannot be parsed`, section[i:])
		}

		// Restore the position of the last element to the default.
		// Current symbol continues the section name and thus the assumption
		// that the previous space was trailing is incorrect.
		endInd = l
	}

	// Make sure that all of the square brackets are closed.
	if unclosedBr != 0 {
		return nil, errors.New("not all square brackets are closed")
	}

	// Return the result not including the trailing spaces.
	return section[:endInd], nil
}
