package parser

// parseLine gets an arbitrary line of INI configuration file
// and tryes to parse it.
func (c *context) parseLine(line []byte) error {
	// Clean the trailing spaces.
	line, l := trimSpaceLeft(line)
	if l == 0 {
		return nil
	}

	// Check what the current line looks like
	// and process appropriately.
	switch line[0] {
	case commentBeg:
		// Omit the comment.
	case sectionBeg:
		// Parse the section and append it to the list of results.
		section, err := c.parseSection(line[1:])
		if err != nil {
			return err
		}
		c.sections = append(c.sections, Section{Name: section})
	default:
		// By default, treat the line as a key-value pair.
		// Add it to the last section that was parsed.
		k, v, err := c.parseKV(line)
		if err != nil {
			return err
		}
		c.sections[len(c.sections)-1].add(k, v)
	}
	return nil
}
