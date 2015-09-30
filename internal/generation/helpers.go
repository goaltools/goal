package generation

import (
	"fmt"
	"path/filepath"
	"text/template"
)

var funcs = template.FuncMap{
	"base":    filepath.Base,
	"dict":    dict,
	"join":    filepath.Join,
	"set":     set,
	"sprintf": fmt.Sprintf,
}

// dict joins a bunch of maps into one and returns it.
func dict(as ...map[string]interface{}) map[string]interface{} {
	m := map[string]interface{}{}
	for i := 0; i < len(as); i++ {
		for k := range as[i] {
			m[k] = as[i][k]
		}
	}
	return m
}

// set gets a key string and a value interface{}
// and returns them as a map.
func set(k string, v interface{}) map[string]interface{} {
	m := map[string]interface{}{
		k: v,
	}
	return m
}
