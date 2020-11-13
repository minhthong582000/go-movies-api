package usecase

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/minhthong582000/go-movies-api/config"
	"github.com/minhthong582000/go-movies-api/domain"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	SecretKey string
	Issuer    string
	UserRepo  domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) domain.AuthUsecase {
	return &authUsecase{
		SecretKey: getSecretKey(),
		Issuer:    "minhthong/go-movies-api",
		UserRepo:  userRepo,
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

func (a *authUsecase) SignIn(credentials domain.Credentials) (token string, err error) {
	existed, err := a.UserRepo.GetByUsername(credentials.Username)
	fmt.Printf("%v\n", existed)
	if existed == (domain.User{}) {
		return "", err
	}

	isValid, err := a.UserRepo.ComparePassword(credentials.Password, existed.Password)
	if isValid == false {
		return "", err
	}

	token, err = a.GenerateToken(existed.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *authUsecase) SignUp(user *domain.User) (token string, err error) {
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	err = a.UserRepo.Store(user)
	if err != nil {
		return "", err
	}

	token, err = a.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func getSecretKey() string {
	secretKey := config.Env("JWT_SECRET")
	if secretKey == "" {
		return "TempKey"
	}

	return secretKey
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
