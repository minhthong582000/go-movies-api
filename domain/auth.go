package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Credentials DTO
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthUsecase interface {
	GenerateToken(username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	SignIn(credentials Credentials) (token string, err error)
	SignUp(user *User) (token string, err error)
}

type AuthMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
}
