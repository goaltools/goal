package output

import (
	//"strings"
	"path/filepath"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/anonx/sunplate/reflect"
)

var funcs = template.FuncMap{
	"export": export,
	"join":   filepath.Join,
}

func parseFunc(arg reflect.Arg) (string, error) {
	t := arg.Type.String()
	return t, nil
}

// export gets a lowercased string such as "int" or "string"
// and transforms it into an exported form, e.g. "Int".
func export(name string) string {
	if len(name) == 0 {
		return name
	}
	r, size := utf8.DecodeRuneInString(name)
	return string(unicode.ToUpper(r)) + name[size:]
}
