// DO NOT EDIT! This file is generated automatically by goctl-pet copygen.
package middleware

import (
	"net/http"

	"github.com/estatie/auth/jwt"
)

type RoleMiddleware struct {
	roles  []string
	handle func(http.HandlerFunc) http.HandlerFunc
}

func NewRoleMiddleware(secret string, roles ...string) *RoleMiddleware {
	return &RoleMiddleware{
		roles:  roles,
		handle: jwt.NewMiddleware(secret).Roles(roles...),
	}
}

func (m *RoleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return m.handle(next)
}
