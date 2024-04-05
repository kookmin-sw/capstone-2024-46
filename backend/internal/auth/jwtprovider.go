package auth

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"private-llm-backend/pkg/errorutil"
	"private-llm-backend/pkg/pointerutil"
)

type JWTProvider interface {
	Generate(claims *JWTClaims, duration time.Duration) (string, error)
	Verify(tokenString string) (*JWTClaims, error)
}

var _ JWTProvider = (*jwtProvider)(nil)

type jwtProvider struct {
	secret []byte
}

func (j *jwtProvider) Generate(claims *JWTClaims, duration time.Duration) (string, error) {
	createTime := time.Now()
	if claims.RegisteredClaims.IssuedAt != nil {
		createTime = claims.RegisteredClaims.IssuedAt.Time
	} else {
		claims.RegisteredClaims.IssuedAt = jwt.NewNumericDate(createTime)
	}
	expireTime := createTime.Add(duration)
	if claims.RegisteredClaims.ExpiresAt == nil {
		claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expireTime)
	}
	claims.JWTVersion = pointerutil.Int(claims.LatestVersion())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		err = errorutil.WithDetail(err, errors.New("failed to sign token"))
		return "", err
	}
	return tokenString, nil
}

func (j *jwtProvider) Verify(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the token's signing algorithm
		if token.Method.Alg() == "HS256" {
			// If using HMAC algorithm, return the secret key for signature verification
			return j.secret, nil
		}
		// If an unexpected signing method is used, return an error
		return nil, errors.New("unexpected signing method")
	})
	if err != nil {
		return nil, errorutil.WithDetail(err, ErrInvalidJWT)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errorutil.Error(ErrInvalidJWT)
	}
	if claims.JWTVersion == nil || claims.LatestVersion() != *claims.JWTVersion {
		return nil, errorutil.Error(ErrInvalidJWT)
	}

	return claims, nil
}

func NewJWTProvider(secretInBase64 string) (JWTProvider, error) {
	secret, err := base64.StdEncoding.DecodeString(secretInBase64)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to decode secret"))
	}
	return &jwtProvider{
		secret: secret,
	}, nil
}
