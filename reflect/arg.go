package reflect

import (
	"go/ast"
	"strings"
)

// Arg is used to describe arguments of functions and fields of structures.
type Arg struct {
	Name string // Name of the argument, e.g. "name" or "age".
	Tag  string // Tag is a field tag that may be presented.
	Type *Type  // Type represents a type of argument.
}

// processFieldList expects an ast FieldList as input parameter.
// The list is transformed into []Arg.
func processFieldList(fields *ast.FieldList) (list []Arg) {
	// Make sure FieldList is not empty.
	if fields == nil {
		return
	}

	// Extract the info we need.
	for _, field := range fields.List {
		t := processField(field)
		if t != nil {
			list = append(list, t...)
		}
	}
	return
}

// processField receives an ast field structure
// and returns a list of  extracted arguments.
func processField(field *ast.Field) (list []Arg) {
	// All names of the same field have the same type.
	t := processType(field.Type)
	if t == nil { // Skip fields that we don't know how to process.
		return
	}

	// Check whether tag is presented.
	tag := ""
	if field.Tag != nil {
		tag = strings.Trim(field.Tag.Value, "\"`") // Remove quote signs from the left & right sides.
	}

	// If there are no names, return without them.
	if len(field.Names) == 0 {
		return []Arg{
			{
				Tag:  tag,
				Type: t,
			},
		}
	}

	// Otherwise, iterate through all the names of this field.
	for _, name := range field.Names {
		// Add current argument to the list.
		list = append(list, Arg{
			Name: name.Name,
			Tag:  tag,
			Type: t,
		})
	}
	return
}
