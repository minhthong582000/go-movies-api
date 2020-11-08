package usecase

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/minhthong582000/go-movies-api/config"
	"github.com/minhthong582000/go-movies-api/domain"
)

type authUsecase struct {
	SecretKey string
	Issuer    string
}

func NewAuthUsecase() domain.AuthUsecase {
	return &authUsecase{
		SecretKey: getSecretKey(),
		Issuer:    "minhthong/go-movies-api",
	}
}

func (a *authUsecase) GenerateToken(username string) (string, error) {
	// Set custom and standard claims
	claims := &domain.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			Issuer:    a.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(a.SecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *authUsecase) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(a.SecretKey), nil
	})
}

func getSecretKey() string {
	secretKey := config.Env("JWT_SECRET")
	if secretKey == "" {
		return "TempKey"
	}

	return secretKey
}
