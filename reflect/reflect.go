// Package reflect is ...
package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/anonx/sunplate/log"
)

// Func is a type that represents information about a function or method.
type Func struct {
	Args     []Arg    // A list of arguments this function receives.
	Comments []string // Comments that are located right above the function declaration.
	Line     int      // Line of code where this function has been found.
	Name     string   // Name of the function, e.g. "Index" or "About".
	Return   []Arg    // A list of arguments the function returns.
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

// processDecl receives an ast declaration, checks whether it is
// of general type and represents a structure.
// If so, it returns a *Struct. Otherwise, nil is returned.
func processDecl(decl ast.Decl) *Struct {
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

// processCommentGroup is a simple function that transforms *ast.CommentGroup
// into a slice of strings.
func processCommentGroup(group *ast.CommentGroup) (list []string) {
	// Make sure comments do exist at all.
	if group == nil {
		return
	}

	// If they are, add them to the list and return it.
	for _, comment := range group.List {
		list = append(list, comment.Text)
	}
	return
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
