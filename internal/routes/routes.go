// Package routes is used for parsing route comments
// and route field tags.
package routes

import (
	"path"
	"reflect"
	"strings"

	"github.com/colegion/goal/internal/log"
	r "github.com/colegion/goal/internal/reflect"
)

const wildcardRoute = "ROUTE"

var methods = map[string]bool{
	"get": true, "head": true, "post": true, "put": true, "delete": true,
	"trace": true, "options": true, "connect": true, "patch": true,
	strings.ToLower(wildcardRoute): true,
}

// Prefixes stores information about route prefixes
// of a controller.
type Prefixes []Route

// Route represents a single route, i.e.
// pattern and an associated method.
type Route struct {
	Pattern, Method string
	Default         bool
}

// ParseRoutes gets an action function and returns all the
// routes that are presented in its comments
// concatenated with the specified prefixes.
func (ps Prefixes) ParseRoutes(f *r.Func) (rs []Route) {
	for i := range f.Comments {
		// Skip comments that do not contain routes.
		m, p, d, ok := parseComment(f.Comments[i])
		if !ok {
			continue
		}

		// Concatenate route with every of the prefix
		// if their methods match.
		for j := range ps {
			// Ignore prefixes which's methods do not match and are not wildcard.
			if ps[j].Method != m && ps[j].Method != wildcardRoute {
				continue
			}

			// Concatenate all other prefixes and add to the list.
			println(m)
			rs = append(rs, Route{
				Method:  m,
				Pattern: path.Join(ps[j].Pattern, p),
				Default: d,
			})
		}
	}
	return
}

// ParseTag gets a tag string and extracts route prefixes out of it.
// If that's not possible, {}, false is returned.
func ParseTag(t string) (ps Prefixes) {
	// Parse prefix for every of the methods.
	st := reflect.StructTag(t)
	for m := range methods {
		// Make sure the prefix is presented.
		v := st.Get("@" + m)
		if v == "" {
			continue
		}

		// Add the prefix to the result list.
		ps = append(ps, Route{
			Pattern: strings.TrimSpace(v),
			Method:  strings.ToUpper(m),
		})
	}
	return
}

// parseComment gets a line comment, parses it, and returns
// route method, pattern, and true.
// If comment doesn't contain a route, "", "", false will be returned.
func parseComment(c string) (method, pattern string, def, ok bool) {
	// Route comments must start with "//@".
	if !strings.HasPrefix(c, "//@") {
		return
	}

	// Make sure the comment contains a correct method.
	// NB: They must be lowecased.
	cs := strings.SplitN(c[3:], " ", 2)
	if _, ok = methods[cs[0]]; !ok {
		log.Warn.Printf(
			`Comment "%s contains incorrect method "%s". Supported ones are %v.`,
			c, cs[0], keys(methods),
		)
		return
	}
	method = strings.ToUpper(cs[0]) // Result will be uppercased.
	ok = true

	// Check whether pattern is empty.
	if len(cs) <= 1 {
		def = true
		return
	}
	pattern = strings.TrimSpace(cs[1])
	return
}

// keys returns keys of the map[string]bool as a slice.
func keys(m map[string]bool) (ks []string) {
	for k := range m {
		ks = append(ks, strings.ToUpper(k))
	}
	return
}
