package run

import (
	"strings"
)

// varList are special variables in tasks that
// should be dinamically replaced by some values.
var varList = map[string]string{
	":EXT": "", // Extension of the generated binary (empty by-default, ".exe" on Win).
}

func replaceVars(s string) string {
	for k, v := range varList {
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
