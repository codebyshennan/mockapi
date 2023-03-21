package mw

import (
	"context"
	"errors"
	"net/http"

	"github.com/codebyshennan/mockapi/internal"
	"github.com/codebyshennan/mockapi/internal/common/h"
)

type MiddlewareProvider struct {
	provider *internal.Provider
}

func NewMiddlewareProvider(p *internal.Provider) (m *MiddlewareProvider) {
	m = &MiddlewareProvider{
		provider: p,
	}

	if p == nil {
		return nil
	}
	return m
}

// Validates the JWT, decodes it and sets the decoded UserID and
// Roles in the request context
func (o MiddlewareProvider) AuthenticateTokenAndSetContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if o.provider.Config.DisableAuth {
			next.ServeHTTP(w, r)
			return
		}

		token := getTokenFromHeader(r)
		if token == "" {
			o.provider.Logger.Infoln("ERR_AUTH_TODO", errors.New("No token found"))
			h.JsonRes(w, http.StatusUnauthorized, map[string]string{
				"code": "ERR_AUTH_TODO",
			}, o.provider.Logger)
			return
		}

		user, err := o.provider.AuthRepo.ValidateAndDecodeJWT(token)
		if err != nil {
			o.provider.Logger.Infoln("ERR_AUTH_TODO", err.Error())
			h.JsonRes(w, http.StatusUnauthorized, map[string]string{
				"code": "ERR_AUTH_TODO",
			}, o.provider.Logger)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "user", user))
		next.ServeHTTP(w, r)
	})
}

func getTokenFromHeader(r *http.Request) (token string) {
	if r.Header["Authorization"] == nil {
		return
	}

	// Retrieve token from header in the form "Bearer ..."
	token = r.Header["Authorization"][0][7:]
	return
}
