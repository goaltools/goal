package reflect

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
)

// Struct is a type that represents information about a specific struct,
// its fields, and comments group.
type Struct struct {
	Comments []string // Comments right above the struct declaration.
	Fields   []Arg    // A list of fields that belong to this struct.
	File     string   // Name of the file where the function is located.
	Name     string   // Name of the struct, e.g. "Application".
}

// processStructDecl ensures that received ast gen declaration
// represents a structure, parses it, and returns.
// If input data is not correct, nil will be returned.
func processStructDecl(decl *ast.GenDecl) *Struct {
	// Make sure it is a type declaration.
	if decl.Tok != token.TYPE {
		return nil
	}

	// Compose a structure and return it.
	for _, spec := range decl.Specs {
		ts, _ := spec.(*ast.TypeSpec)              // TypeSpec is the only possible value, so ignoring second arg.
		s := processTypeSpec(ts)                   // Composing a structure.
		s.Comments = processCommentGroup(decl.Doc) // Adding comments block.
		return s                                   // There is just one spec in the Specs anyway, so returning.
	}

	// No specs have been found.
	return nil
}

// processImportDecl returns a name - value pairs of imports if
// genDecl's Tok is IMPORT.
// Otherwise, nil is returned.
func processImportDecl(decl *ast.GenDecl) map[string]string {
	// Make sure it is an import declaration.
	if decl.Tok != token.IMPORT {
		return nil
	}

	// Generate a map and return it.
	list := make(map[string]string, len(decl.Specs))
	for _, spec := range decl.Specs {
		is, _ := spec.(*ast.ImportSpec) // ImportSpec is the only possible value, so ignoring second arg.
		k, v := processImportSpec(is)
		list[k] = v
	}
	return list
}

// processTypeSpec expects ast type spec as input parameter
// that is transformed into *Struct representation and returned.
func processTypeSpec(spec *ast.TypeSpec) *Struct {
	// Make sure it is a structure type. Return nil if not.
	structType, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil
	}

	// Compose a structure and return it.
	return &Struct{
		Fields: processFieldList(structType.Fields),
		Name:   spec.Name.Name,
	}
}

// processImportSpec gets ast import spec as input parameter
// and transforms it into a pair of import's name and its value.
func processImportSpec(spec *ast.ImportSpec) (name, value string) {
	// Remove quote signes, etc. from the left and right sides of import.
	value = strings.Trim(spec.Path.Value, "\"`")

	// By-default, use the last (base) part of import as a name.
	name = filepath.Base(value)

	// However, if name value is presented in spec, use it instead.
	if n := spec.Name; n != nil {
		name = n.Name
	}

	return
}
