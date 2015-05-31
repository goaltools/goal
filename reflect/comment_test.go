package reflect

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestProcessCommentGroup_EmptyGroup(t *testing.T) {
	if c := processCommentGroup(nil); len(c) != 0 {
		t.Errorf("Zero length slice expected. Gor %#v instead.", c)
	}
}

func TestProcessCommentGroup(t *testing.T) {
	c := processCommentGroup(&ast.CommentGroup{
		List: []*ast.Comment{
			{
				Text: "// This is line 1",
			},
			{
				Text: "// This is line 2",
			},
			{
				Text: "// This is line 3",
			},
		},
	})
	expRes := []string{
		"// This is line 1", "// This is line 2", "// This is line 3",
	}
	if !reflect.DeepEqual(c, expRes) {
		t.Errorf("Incorrect result of processCommentGroup. Expected %#v, got %#v.", expRes, c)
	}
}
