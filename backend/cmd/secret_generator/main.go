package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"private-llm-backend/internal/auth"
	"private-llm-backend/internal/config"
)

// generateSecretKey 는 HS256 JWT 에 사용될 시크릿 키(32byte) 를 랜덤으로 생성 후 base64 encode 하여 리턴합니다.
func generateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	c, _, err := config.Load(context.Background(), "./configs/", env)
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}
	j, err := auth.NewJWTProvider(*c.JWTSecret)
	if err != nil {
		log.Fatalf("failed to create JWTProvider: %+v", err)
	}

	oneYear := time.Hour * 24 * 365
	jwtToken, err := j.Generate(
		&auth.JWTClaims{
			UserID: "1",
			Email:  "admin@blast-team.com",
		},
		oneYear,
	)
	if err != nil {
		log.Fatalf("failed to generate JWT: %+v", err)
	}
	fmt.Printf("User JWT:\n%s\n\n", jwtToken)
}
