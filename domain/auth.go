package domain

import (
	"github.com/golang-jwt/jwt"
)

// Represents the data that is used to check against Google
// for authentication
type GoogleLoginPostData struct {
	ClientId   string `json:"clientId"`
	Credential string `json:"credential"`
}

// Represents the payload encrypted into a JWT token
type JWTClaims struct {
	UserID string
	jwt.StandardClaims
}

// Represents the return result for authentication
type AuthRes struct {
	Token     string `json:"token"`
	UserModel `json:"user"`
}

type IAuthRepo interface {
	GoogleLogin(d *GoogleLoginPostData) (*AuthRes, error)
	ValidateAndDecodeJWT(token string) (*UserModel, error)
}
