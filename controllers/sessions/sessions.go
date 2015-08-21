// Package sessions implements COOKIE based sessions.
package sessions

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var (
	// CookieName is a name of the cookie where session data
	// is stored.
	CookieName = "_SESSION"

	// CookieSecure should be set to true in order to forbid transmitting
	// cookie over unencrypted connection.
	CookieSecure = false

	// HTTPOnly forbids client side scripts access session cookies.
	HTTPOnly = true

	// Domain of the session cookie.
	Domain = ""

	hashKey = securecookie.GenerateRandomKey(64)
	s       = securecookie.New(hashKey, nil)
)

// Sessions is a controller that makes Session field
// available for your actions when you're using this
// controller as a parent.
type Sessions struct {
	Session map[string]string
}

// Before is a magic actions of Sessions controller.
func (c *Sessions) Before() http.Handler {
	return nil
}

// Initially is a magic method that gets session info from a request
// and initializes Session field.
func (c *Sessions) Initially(w http.ResponseWriter, r *http.Request) bool {
	c.Session = map[string]string{}
	if cookie, err := r.Cookie(CookieName); err == nil {
		s.Decode(CookieName, cookie.Value, &c.Session)
	}
	return false
}

// Finally is a magic method that will be executed at the very end of request
// life cycle and is responsible for creating a signed cookie with session info.
func (c *Sessions) Finally(w http.ResponseWriter, r *http.Request) bool {
	if encoded, err := s.Encode(CookieName, c.Session); err == nil {
		cookie := &http.Cookie{
			Name:     CookieName,
			Value:    encoded,
			Domain:   Domain,
			HttpOnly: HTTPOnly,
			Path:     "/",
			Secure:   CookieSecure,
		}
		http.SetCookie(w, cookie)
	}
	return false
}

// Init is a function that is used for initialization of
// Sessions controller.
// It gets a key argument that will be used for authentication
// of cookie value using HMAC.
func Init(key []byte) {
	hashKey = key
	s = securecookie.New(key, hashKey)
}
