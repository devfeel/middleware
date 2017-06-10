package authz

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/devfeel/dotweb"
)

// AuthzMiddleware represents Casbin Authorization Middleware.
type AuthzMiddleware struct {
	dotweb.BaseMiddlware
	Enforcer *casbin.Enforcer
}

// Handle filters the HTTP request.
func (m *AuthzMiddleware) Handle(ctx dotweb.Context) error {
	if !CheckPermission(m.Enforcer, ctx.Request().Request) {
		ctx.Response().SetStatusCode(403)
		return nil
	}

	m.Next(ctx)
	return nil
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func CheckPermission(e *casbin.Enforcer, r *http.Request) bool {
	user := GetUserName(r)
	method := r.Method
	path := r.URL.Path
	return e.Enforce(user, path, method)
}


// Middleware creates an Authorization Middleware
func Middleware(e *casbin.Enforcer) *AuthzMiddleware {
	return &AuthzMiddleware{Enforcer: e}
}
