package h

import (
	"errors"
	"net/http"

	"github.com/codebyshennan/mockapi/domain"
)

type contextKey string

// Implements String interface
func (c contextKey) String() string {
	return string(c)
}

// Keys in golang context
const (
	UserIDContextKey contextKey = contextKey("user")
)

// Getter for user context
func GetUserIdContext(r *http.Request) (domain.UserModel, error) {
	if user, ok := r.Context().Value("user").(*domain.UserModel); ok {
		return *user, nil
	} else {
		return domain.UserModel{}, errors.New("ERR_AUTH_TODO")
	}
}
