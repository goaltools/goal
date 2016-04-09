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

var (
	supportedMethods = map[string]bool{
		"get": true, "head": true, "post": true, "put": true, "delete": true,
		"trace": true, "options": true, "connect": true, "patch": true,
		strings.ToLower(wildcardRoute): true,
	}
	realMethodsList = []string{
		"GET", "HEAD", "POST", "PUT", "DELETE", "TRACE", "OPTIONS", "CONNECT", "PATCH",
	}
	routePartsSep = map[byte]bool{
		' ': true, '\t': true,
	}
)

// Prefixes stores information about route prefixes
// of a controller.
type Prefixes []Route

// NewPrefixes allocates and returns a new Prefixes object.
func NewPrefixes() Prefixes {
	return Prefixes{
		{Method: wildcardRoute},
	}
}

// Route represents a single route, i.e.
// pattern and an associated method.
type Route struct {
	Pattern, Method, HandlerName, Label string
}

// ParseRoutes gets a controller name and an action function and returns all the
// routes that are presented in its comments concatenated with the specified prefixes.
// If method is specified however pattern isn't, controller/action combination
// is used as a pattern.
func (ps Prefixes) ParseRoutes(controller string, f *r.Func) (rs []Route) {
	for i := range f.Comments {
		// Skip comments that do not contain routes.
		method, pattern, label, ok := parseComment(f.Comments[i])
		if !ok {
			continue
		}

		// If no pattern specified, use controller's and action's names.
		if pattern == "" {
			pattern = path.Join("/", controller, f.Name)
		}

		// Concatenate route with every of the prefixes
		// if their methods match.
		for j := range ps {
			// Ignore prefixes which's methods do not match and they are not wildcard.
			if ps[j].Method != method && ps[j].Method != wildcardRoute {
				continue
			}

			// Concatenate all other prefixes and add to the list.
			ms := realMethods(method)
			for k := range ms {
				r := Route{
					Method:      ms[k],
					Pattern:     path.Join(ps[j].Pattern, pattern),
					HandlerName: controller + "." + f.Name,
					Label:       label,
				}
				log.Trace.Printf(
					`Detected route "%s" "%s" "%s" ("%s")`, r.Method, r.Pattern, r.HandlerName, r.Label,
				)
				rs = append(rs, r)
			}
		}
	}
	return
}

// realMethods gets an HTTP method and returns it as is if the method is real.
// If it is a wildcard pseudo method, all available methods are returned.
func realMethods(m string) []string {
	if m == wildcardRoute {
		return realMethodsList
	}
	return []string{m}
}

// ParseTag gets a tag string and extracts route prefixes out of it.
// If that's not possible, {}, false is returned.
func ParseTag(t string) (ps Prefixes) {
	// Parse prefix for every of the methods.
	st := reflect.StructTag(t)
	for m := range supportedMethods {
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
// If comment doesn't contain a route, "", "", "", false, false will be returned.
func parseComment(c string) (method, pattern, label string, ok bool) {
	// Route comments must start with "//@".
	if !strings.HasPrefix(c, "//@") {
		return
	}

	// Make sure the comment contains a correct method.
	// NB: They must be lowecased.
	cs := splitN(c[3:], 3)
	if _, ok = supportedMethods[cs[0]]; !ok {
		log.Warn.Printf(
			`Comment "%s contains incorrect method "%s". Supported ones are %v.`,
			c, cs[0], keys(supportedMethods),
		)
		return
	}
	method = strings.ToUpper(cs[0]) // Result will be uppercased.
	ok = true

	// Set pattern and label of the route.
	pattern = cs[1]
	label = cs[2]
	return
}

// splitN gets a string and splits it into multiple groups.
func splitN(s string, n int) []string {
	res := make([]string, n)
	sep := true
	j := -1

	// Iterate over characters of the input string.
	for i := 0; i < len(s); i++ {
		// If the current character is a separator,
		// mark it and ignore.
		if routePartsSep[s[i]] {
			sep = true
			continue
		}

		// Current character is not a separator.
		// If the previous one was a separator, increment elements index.
		if sep {
			sep = false
			j++
		}

		// Check whether the limit of results has been reached.
		if j >= len(res) {
			break
		}

		// Add the current symbol to the result.
		res[j] += string(s[i])
	}
	return res
}

// keys returns keys of the map[string]bool as a slice.
func keys(m map[string]bool) (ks []string) {
	for k := range m {
		ks = append(ks, strings.ToUpper(k))
	}
	return
}
