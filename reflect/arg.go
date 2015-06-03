package reflect

import (
	"go/ast"
)

// Args is a type that is used for representation of an arguments list.
type Args []Arg

// Arg is used to describe arguments of functions and fields of structures.
type Arg struct {
	Name string // Name of the argument, e.g. "name" or "age".
	Tag  string // Tag is a field tag that may be presented.
	Type *Type  // Type represents a type of argument.
}

// Filter returns a list of functions from members of a list
// fulfilling a condition given by the fn argument.
func (as Args) Filter(fn func(f *Arg) bool) (res Args) {
	for _, v := range as {
		if fn(&v) {
			res = append(res, v)
		}
	}
	return res
}

// processFieldList expects an ast FieldList as input parameter.
// The list is transformed into Args.
func processFieldList(fields *ast.FieldList) (list Args) {
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
func processField(field *ast.Field) (list Args) {
	// All names of the same field have the same type.
	t := processType(field.Type)
	if t == nil { // Skip fields that we don't know how to process.
		return
	}

	// Check whether tag is presented.
	tag := ""
	if field.Tag != nil {
		// Remove quote signs from the left & right sides.
		tag = field.Tag.Value[1 : len(field.Tag.Value)-1]
	}

	// If there are no names, return without them.
	if len(field.Names) == 0 {
		return Args{
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
