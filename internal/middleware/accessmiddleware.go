// DO NOT EDIT! This file is generated automatically by goctl-pet copygen.
package middleware

import (
	"net/http"

	"github.com/estatie/auth/jwt"
)

type AccessMiddleware struct {
	auth *jwt.Middleware
}

func NewAccessMiddleware(secret string) *AccessMiddleware {
	return &AccessMiddleware{
		auth: jwt.NewMiddleware(secret),
	}
}

func (m *AccessMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	// verify first, then access token check
	return m.auth.Access(next)
}
