package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/minhthong582000/go-movies-api/domain"
)

// AuthMiddleware represent the data-struct for middleware
type authMiddleware struct {
	AuthUsecase domain.AuthUsecase
}

// NewAuthMiddleware initialize the middleware
func NewAuthMiddleware(au domain.AuthUsecase) domain.AuthMiddleware {
	return &authMiddleware{
		AuthUsecase: au,
	}
}

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func (au *authMiddleware) AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromHeader(c.GetHeader("Authorization"))
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := au.AuthUsecase.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["username"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func extractTokenFromHeader(headerString string) string {
	const BEARER_SCHEMA = "Bearer"
	tokenString := strings.Split(headerString, " ")
	if tokenString[0] != BEARER_SCHEMA {
		return ""
	}

	return tokenString[1]
}
