package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	JWTVersion *int `json:"v"`

	UserID string `json:"uid"`
	Email  string `json:"e,omitempty"`
}

// LatestVersion returns the latest version of JWTClaims.
// if you add a new attributes in JWTClaims, you should update this value.
func (c *JWTClaims) LatestVersion() int {
	return 1
}
