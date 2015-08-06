package output

import (
	"reflect"
	"testing"
)

func TestDict(t *testing.T) {
	if r := dict(
		set("name", "James"),
		set("age", 8),
		set("country", "Lemonland"),
		set("verified", true),
	); !reflect.DeepEqual(r, map[string]interface{}{
		"name":     "James",
		"age":      8,
		"country":  "Lemonland",
		"verified": true,
	}) {
		t.Errorf("Incorrect dict result: %v.", r)
	}
}
