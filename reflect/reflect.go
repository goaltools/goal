// Package reflect is a wrapper for go/ast, go/token, and go/parser packages.
// It is used to get information about functions, methods, structures, and
// imports of go files in a specific directory.
package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/anonx/sunplate/log"
)

// Imports is a map of import paths in the following format:
//	- Filename:
//		- Import name:
//			- Import value
type Imports map[string]map[string]string

// Methods is a map of functions with receiver in the following format:
//	- Name of a struct:
//		- Methods
type Methods map[string]Funcs

// Package is a type that combines declarations
// of functions, types, and structs of a single go package.
type Package struct {
	Funcs   Funcs   // A list of functions of the package.
	Imports Imports // Imports of this package grouped by files.
	Methods Methods // Struct names and their Methods (functions with receivers).
	Name    string  // Name of the package, e.g. "controllers".
	Structs Structs // A list of struct types of the package.
}

// Value checks whether requested import name exists in
// requested file. If so, import value and true are returned.
// Otherwise, empty string and false will be the results.
func (i Imports) Value(file, name string) (string, bool) {
	// Check whether such file exists.
	f, ok := i[file]
	if !ok {
		return "", false
	}

	// Make sure requested name does exist.
	v, ok := f[name]
	if !ok {
		return "", false
	}
	return v, true
}

// Name checks whether an import that ends with a requested value exists
// in the requested file. If so, it's name and true are returned.
// Otherwise, empty string and false will be the results.
// Imports are ensured to end with the value rather than be equal to it
// for compatability with "vendor" package manager.
func (i Imports) Name(file, value string) (string, bool) {
	// Check whether such file exists.
	f, ok := i[file]
	if !ok {
		return "", false
	}

	// Make sure requested import does exist.
	for name := range f {
		if strings.HasSuffix(f[name], value) {
			return name, true
		}
	}
	return "", false
}

// ParseDir expects a path to directory with a go package
// that is parsed and returned in a form of *Package.
func ParseDir(path string) *Package {
	fset := token.NewFileSet() // Positions are relative to fset.
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Error.Panic(err)
	}

	// Just one package per directory is allowed.
	// So, receiving it.
	var pkg *ast.Package
	for _, v := range pkgs {
		pkg = v
		break
	}

	// Iterating through files of the package and combining all declarations
	// into single Package struct.
	p := &Package{
		Imports: map[string]map[string]string{},
		Methods: map[string]Funcs{},
		Name:    pkg.Name,
	}
	for name, file := range pkg.Files {
		// Extract functions, methods, sructures, and imports from file declarations.
		fs, ms, ss, is := processDecls(file.Decls, name)

		// Add functions to the list.
		if len(fs) > 0 {
			p.Funcs = append(p.Funcs, fs...)
		}

		// Attach methods.
		if len(ms) > 0 {
			p.Methods = joinMethods(p.Methods, ms)
		}

		// Add structures to the package.
		if len(ss) > 0 {
			p.Structs = append(p.Structs, ss...)
		}

		// Add imports of the current file.
		p.Imports[name] = is
	}
	return p
}

// processDecls expects a list of declarations as an input
// parameter. It will be parsed, splitted into functions,
// methods, and structs and returned.
func processDecls(decls []ast.Decl, file string) (fs Funcs, ms Methods, ss Structs, is map[string]string) {
	for _, decl := range decls {
		// Try to process the declaration as a function.
		var f *Func
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			f = processFuncDecl(funcDecl)
		}
		if f != nil { // If the decl really was a func declaration.
			f.File = file      // Set name of the file we are processing.
			if f.Recv == nil { // If the function has no receiver.
				// Add the processed decl to the list of functions.
				fs = append(fs, *f)
				continue
			}
			// Otherwise, add it to the list of methods.
			ms = joinMethods(ms, Methods{f.Recv.Type.Name: Funcs{*f}})
			continue
		}

		// It is likely a GenDecl.
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			// Try to process the GenDecl as a structure.
			s := processStructDecl(genDecl)
			if s != nil {
				s.File = file // Set name of the file we are processing.

				// Add the structure to the list.
				ss = append(ss, *s)
				continue
			}

			// Try to process the GenDecl as an import.
			imp := processImportDecl(genDecl)
			if imp != nil {
				// Add the imports to the map.
				is = joinMaps(is, imp)
				continue
			}
		}
	}
	return
}

// joinMaps adds addition map[string]string to the base one.
// If there are key collisions, addition argument's values
// have higher priority.
// NOTE: this function has side effects, base is modified.
func joinMaps(base, addition map[string]string) map[string]string {
	// Make sure base map is initialized.
	if base == nil {
		base = map[string]string{}
	}

	// Join two maps and return the result.
	for k, v := range addition {
		base[k] = v
	}
	return base
}

// joinMethods adds addition Methods to the base one.
// If there are key collisions, addition argument's values
// have higher priority.
// NOTE: this function has side effects, base is modified.
func joinMethods(base, addition Methods) Methods {
	// Make sure base Methods structure is initialized.
	if base == nil {
		base = Methods{}
	}

	// Join two Methods structures and return the result.
	for k, v := range addition {
		base[k] = append(base[k], v...)
	}
	return base
}
