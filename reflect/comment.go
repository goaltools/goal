package reflect

import (
	"go/ast"
)

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
