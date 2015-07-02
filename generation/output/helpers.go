package output

import (
	"path/filepath"
	"text/template"
)

var funcs = template.FuncMap{
	"join": filepath.Join,
}
