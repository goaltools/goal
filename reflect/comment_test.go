package reflect

import (
	"go/ast"
	"reflect"
	"strings"
	"testing"
)

func TestCommentsFilter(t *testing.T) {
	t1 := Comments{
		"Comment1",
		"Comment2",
		"Comment12",
	}
	expRes := Comments{
		"Comment2",
		"Comment12",
	}
	r := t1.Filter(func(s string) bool {
		return true
	})
	if !reflect.DeepEqual(t1, r) {
		t.Errorf("Incorrect Filter result. Expected %#v, got %#v.", t1, r)
	}

	r = t1.Filter(func(s string) bool {
		if strings.HasSuffix(s, "2") {
			return true
		}
		return false
	})
	if !reflect.DeepEqual(expRes, r) {
		t.Errorf("Incorrect Filter result. Expected %#v, got %#v.", expRes, r)
	}
}

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
	expRes := Comments{
		"// This is line 1", "// This is line 2", "// This is line 3",
	}
	if !reflect.DeepEqual(c, expRes) {
		t.Errorf("Incorrect result of processCommentGroup. Expected %#v, got %#v.", expRes, c)
	}
}
