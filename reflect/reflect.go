// Package reflect is ...
package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/anonx/sunplate/log"
)

// Package is a type that is
type Package struct {
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
				_ = funcDecl
			}
			//ast.Print(fset, decl)
		}
		ast.Print(fset, file.Comments)
	}
	ast.Print(fset, pkgs)
}
