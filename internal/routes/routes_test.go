package routes

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	r "github.com/goaltools/goal/internal/reflect"
)

func TestParseRoutes(t *testing.T) {
	ps := Prefixes{
		{
			Method:  "ROUTE",
			Pattern: "/",
		},
		{
			Method:  "GET",
			Pattern: "method/get/",
		},
		{
			Method:  "POST",
			Pattern: "/method/post",
		},
	}
	res := ps.ParseRoutes("Default", &r.Func{
		Comments: []string{
			"// Index is a simple action.",
			"// It automatically bind your action //@post nothing.",
			"//@route stuff",
			"//@get   ",
			"//post /index/page1",
			"//@post /index/page",
			"//@xxx /hey",
			"//Some stuff is here...",
		},
		Name: "Index",
	})
	exp := []Route{
		{ // Begin of method wildcardRoute.
			Method:  "GET",
			Pattern: "/stuff",
		},
		{
			Method:  "HEAD",
			Pattern: "/stuff",
		},
		{
			Method:  "POST",
			Pattern: "/stuff",
		},
		{
			Method:  "PUT",
			Pattern: "/stuff",
		},
		{
			Method:  "DELETE",
			Pattern: "/stuff",
		},
		{
			Method:  "TRACE",
			Pattern: "/stuff",
		},
		{
			Method:  "OPTIONS",
			Pattern: "/stuff",
		},
		{
			Method:  "CONNECT",
			Pattern: "/stuff",
		},
		{
			Method:  "PATCH",
			Pattern: "/stuff",
		}, // End of method wildcardRoute.
		{
			Method:  "GET",
			Pattern: "/Default/Index",
		},
		{
			Method:  "GET",
			Pattern: "method/get/Default/Index",
		},
		{
			Method:  "POST",
			Pattern: "/index/page",
		},
		{
			Method:  "POST",
			Pattern: "/method/post/index/page",
		},
	}
	if !equalPrefixes(res, exp) {
		t.Errorf("Expected %v.\nGot %v.", exp, res)
	}
}

func TestParseTag(t *testing.T) {
	for _, v := range []struct {
		tag string
		ps  Prefixes
	}{
		{
			tag: "",
		},
		{
			tag: `sql:"test"`,
		},
		{
			tag: `@trace:"  /tc/xxx   "`,
			ps: Prefixes{
				{Pattern: "/tc/xxx", Method: "TRACE"},
			},
		},
		{
			tag: `@route:"/rt" @get:"/gt" @post:"/ps"`,
			ps: Prefixes{
				{Pattern: "/rt", Method: "ROUTE"},
				{Pattern: "/gt", Method: "GET"},
				{Pattern: "/ps", Method: "POST"},
			},
		},
	} {
		ps := ParseTag(v.tag)
		if !equalPrefixes(ps, v.ps) {
			t.Errorf(
				`"%s": Expected prefix "%v", got prefix "%v".`,
				v.tag,
				v.ps, ps,
			)
		}
	}
}

// equalPrefixes returns true if input prefixes are equal,
// and false otherwise.
func equalPrefixes(p1, p2 Prefixes) bool {
	if len(p1) != len(p2) {
		return false
	}
	sort.Sort(byMethodAndPattern(p1))
	sort.Sort(byMethodAndPattern(p2))
	for i := range p1 {
		if p1[i].Method != p2[i].Method || p1[i].Pattern != p2[i].Pattern {
			fmt.Printf(
				`Prefix "%s" "%s" not equal to "%s" "%s".`,
				p1[i].Method, p1[i].Pattern,
				p2[i].Method, p2[i].Pattern,
			)
			return false
		}
	}
	return true
}

type byMethodAndPattern Prefixes

func (p byMethodAndPattern) Len() int      { return len(p) }
func (p byMethodAndPattern) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byMethodAndPattern) Less(i, j int) bool {
	return p[i].Method < p[j].Method && p[i].Pattern < p[j].Pattern
}

func TestParseComment(t *testing.T) {
	for _, v := range []struct {
		comment, method, pattern, label string
		ok                              bool
	}{
		{
			comment: "// @route",
			ok:      false,
		},
		{
			comment: "//@stuff",
			ok:      false,
		},
		{
			comment: "//@PUT",
			ok:      false,
		},
		{
			comment: "//@get",
			method:  "GET",
			ok:      true,
		},
		{
			comment: "//@post /",
			pattern: "/",
			method:  "POST",
			ok:      true,
		},
		{
			comment: "//@delete  \t    /user/xxx    \r\n",
			pattern: "/user/xxx",
			method:  "DELETE",
			ok:      true,
		},
	} {
		if m, p, l, ok := parseComment(v.comment); m != v.method || p != v.pattern || ok != v.ok {
			t.Errorf(
				`"%s": Expected "%v" method "%s" "%s" ("%s"), got "%v" method "%s" "%s" ("%s").`,
				v.comment,
				v.ok, v.method, v.pattern, v.label,
				ok, m, p, l,
			)
		}
	}
}

func TestNewPrefixes(t *testing.T) {
	ps := NewPrefixes()
	if len(ps) != 1 || ps[0].Method != wildcardRoute {
		t.Fail()
	}
}

func TestSplitN(t *testing.T) {
	for _, v := range []struct {
		s   string
		n   int
		exp []string
	}{
		{"", 0, []string{}},
		{"", 1, []string{""}},
		{"one two", 1, []string{"one"}},
		{"  one   \t  two    \t  ", 2, []string{"one", "two"}},
		{"one   \t  two    \t  three", 2, []string{"one", "two"}},
		{"     one     ", 5, []string{"one", "", "", "", ""}},
	} {
		if res := splitN(v.s, v.n); !reflect.DeepEqual(res, v.exp) {
			t.Errorf(`"%s" (%d): Expected %#v, got %#v.`, v.s, v.n, v.exp, res)
		}
	}
}
