package main

import (
	"fmt"

	_authHandler "github.com/minhthong582000/go-movies-api/auth/delivery/http"
	_authHttpMiddleware "github.com/minhthong582000/go-movies-api/auth/delivery/http/middleware"
	_authUsecase "github.com/minhthong582000/go-movies-api/auth/usecase"
	"github.com/minhthong582000/go-movies-api/config"
	"github.com/minhthong582000/go-movies-api/domain"
	"github.com/minhthong582000/go-movies-api/middleware"
	_movieHandler "github.com/minhthong582000/go-movies-api/movie/delivery/http"
	_gormMovieRepo "github.com/minhthong582000/go-movies-api/movie/repository/gorm"
	_movieUsecase "github.com/minhthong582000/go-movies-api/movie/usecase"

	_userRepo "github.com/minhthong582000/go-movies-api/user/repository/gorm"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	middleware.SetupLogOutput()

	db, err := gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	db.AutoMigrate(&domain.Movie{}, &domain.Director{}, &domain.Actor{}, &domain.Gendres{}, &domain.TrailerLink{}, &domain.User{})

	// Init user repo
	userRepo := _userRepo.NewGormUserRepository(db)

	// Init auth http middleware
	authUse := _authUsecase.NewAuthUsecase(userRepo)
	authHTTPMdw := _authHttpMiddleware.NewAuthMiddleware(authUse)
	_authHandler.NewAuthHandler(r, authUse)

	//
	// PROTECTED ROUTERS
	//
	r.Use(authHTTPMdw.AuthorizeJWT())

	// Init movie handler
	movieRepo := _gormMovieRepo.NewGormMovieRepository(db)
	movieUse := _movieUsecase.NewMovieUsecase(movieRepo)
	_movieHandler.NewMovieHandler(r, movieUse)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
