// Package handlers is generated automatically by "goal generate handlers" tool.
// Please, do not edit it manually.
package handlers

import (
	"net/http"

	contr "github.com/colegion/contrib/controllers/global"

	"github.com/colegion/goal/strconv"
)

// Global is an instance of tGlobal that is automatically generated from
// Global controller being found at "github.com/colegion/contrib/controllers/global/global.go",
// and contains methods to be used as handler functions.
//
// Global is a controller that provides a registry for
// global request variables.
var Global tGlobal

// tGlobal is a type with handler methods of Global controller.
type tGlobal struct {
}

// newC allocates (github.com/colegion/contrib/controllers/global).Global controller,
// initializes its parents and returns it.
func (t tGlobal) newC(w http.ResponseWriter, r *http.Request, ctr, act string) *contr.Global {
	// Allocate a new controller. Set values of special fields, if necessary.
	c := &contr.Global{

		CurrentAction: act,

		CurrentController: ctr,
	}

	return c
}

// before is a method that is started by every handler function at the very beginning
// of their execution phase no matter what. If before returns non-nil result,
// no action methods will be started.
func (t tGlobal) before(c *contr.Global, w http.ResponseWriter, r *http.Request) http.Handler {

	// Call special Before action of (github.com/colegion/contrib/controllers/global).Global.
	if res := c.Before(); res != nil {
		return res
	}

	return nil
}

// after is a method that is started by every handler function at the very end
// of their execution phase no matter what.
func (t tGlobal) after(c *contr.Global, w http.ResponseWriter, r *http.Request) (res http.Handler) {

	return
}

func init() {
	_ = strconv.MeaningOfLife
}
