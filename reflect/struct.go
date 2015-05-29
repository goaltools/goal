package reflect

import (
	"go/ast"
)

// Struct is a type that represents information about a specific struct,
// its methods, etc.
type Struct struct {
	Comments      []string // Comments right above the struct declaration.
	EmbeddedTypes []Struct // Types that are embedded into the struct.
	Fields        []Arg    // A list of fields that belong to this struct.
	File          string   // Name of the file where the struct is located.
	Import        string   // Import path of the struct, e.g. "github.com/anonx/sunplate".
	Methods       []Func   // A list of methods this struct has.
	Name          string   // Name of the struct, e.g. "Application".
	Package       string   // Package name, for example "controllers".
}

// processStruct is a function that extracts information
// from *ast.StructType and returns it in a format that is
// used by this reflect package.
func processStruct(struc *ast.StructType) *Struct {
	return nil
}
