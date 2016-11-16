// +build windows

package importpath

import (
	"strings"
)

// parent detects whether the input path is inside the
// input root directory.
// Paths are case-insensetive. Both are expected to be
// cleaned in advance.
func parent(root, path string) bool {
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(root))
}
