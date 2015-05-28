// Package reflect is ...
package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/anonx/sunplate/log"
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

// Func is a type that represents information about a function or method.
type Func struct {
	Args     []Arg    // A list of arguments this function receives.
	Comments []string // Comments that are located right above the function declaration.
	Line     int      // Line of code where this function has been found.
	Name     string   // Name of the function, e.g. "Index" or "About".
	Return   []Arg    // A list of arguments the function returns.
}

// Arg is used to describe arguments of methods and fields of structures.
type Arg struct {
	Import string // If the argument is of an imported type, this is the import path.
	Name   string // Name of the argument, e.g. "name" or "age".
	Tag    string // Tag is a field tag that may be presented.
	Type   *Type  // Type represents a type of argument.
}

// ParseDir ...
func ParseDir(path string) {
	fset := token.NewFileSet() // Positions are relative to fset.
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Error.Panic(err)
	}

	var pkg *ast.Package
	for k, v := range pkgs {
		log.Error.Println(k)
		pkg = v
	}

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				log.Info.Println(getFuncName(funcDecl))
			}
			//ast.Print(fset, decl)
		}
		ast.Print(fset, file.Comments)
	}
	ast.Print(fset, pkgs)
}

// processStruct receives an ast declaration, checks whether it is
// of general type and represents a structure.
// If so, it returns a *Struct. Otherwise, nil is returned.
func processStruct(decl ast.Decl) *Struct {
	// Make sure the type is of general type.
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil
	}

	// Make sure it has a "type" token.
	if genDecl.Tok != token.TYPE {
		return nil
	}

	// Try to parse fields and embedded types of the struct.
	for _, spec := range genDecl.Specs {
		_ = spec
	}

	return nil
}

// processField receives an ast field structure
func processField(field *ast.Field) *Arg {
	return nil
}

// getFuncName returns a name for this func or method declaration.
// e.g. "(*Application).SayHello" for a method, "SayHello" for a func.
func getFuncName(funcDecl *ast.FuncDecl) string {
	prefix := ""
	if funcDecl.Recv != nil {
		recvType := funcDecl.Recv.List[0].Type
		if recvStarType, ok := recvType.(*ast.StarExpr); ok {
			prefix = "(*" + recvStarType.X.(*ast.Ident).Name + ")"
		} else {
			prefix = recvType.(*ast.Ident).Name
		}
		prefix += "."
	}
	return prefix + funcDecl.Name.Name
}
