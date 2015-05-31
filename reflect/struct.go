package reflect

import (
	"go/ast"
	"go/token"
)

// Struct is a type that represents information about a specific struct,
// its methods, etc.
type Struct struct {
	Comments []string // Comments right above the struct declaration.
	Fields   []Arg    // A list of fields that belong to this struct.
	Name     string   // Name of the struct, e.g. "Application".
}

// processStruct is a function that extracts information
// about structure from *ast.GenDecl.
// If input argument does not represent TypeStruct, nil will be returned.
func processStruct(genDecl *ast.GenDecl) *Struct {
	// Make sure it is a type declaration. If not, return nil.
	if genDecl.Tok != token.TYPE {
		return nil
	}
	return nil
}

// processStructTypeSpec expects ast type spec as input parameter
// that is transformed into *Struct representation and returned.
func processStructTypeSpec(spec *ast.TypeSpec) *Struct {
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
