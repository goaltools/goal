package reflect

import (
	"go/ast"
)

// Comments is a type that is used for representation of a comments list.
type Comments []string

// Filter returns a list of comments from members of a list
// fulfilling a condition given by the fn argument.
func (cs Comments) Filter(fn func(s string) bool) (res Comments) {
	for _, v := range cs {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

// processCommentGroup is a simple function that transforms *ast.CommentGroup
// into a slice of strings.
func processCommentGroup(group *ast.CommentGroup) (list Comments) {
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
