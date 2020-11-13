package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/minhthong582000/go-movies-api/domain"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, au domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: au,
	}

	moviesRouter := r.Group("/auth")
	{
		moviesRouter.POST("/signin", handler.SignIn)
		moviesRouter.POST("/signup", handler.SignUp)
	}
}

func (ah *AuthHandler) SignIn(c *gin.Context) {
	var credentials domain.Credentials

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	token, err := ah.AuthUsecase.SignIn(credentials)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "username": credentials.Username})
}

func (ah *AuthHandler) SignUp(c *gin.Context) {
	var newUser domain.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Validate request obj
	if ok, err := isRequestValid(&newUser); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := ah.AuthUsecase.SignUp(&newUser)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "username": newUser.Username})
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
