package main

import (
	"fmt"

	_authHttpMiddleware "github.com/minhthong582000/go-movies-api/auth/delivery/http/middleware"
	_authUsecase "github.com/minhthong582000/go-movies-api/auth/usecase"
	"github.com/minhthong582000/go-movies-api/config"
	"github.com/minhthong582000/go-movies-api/domain"
	"github.com/minhthong582000/go-movies-api/middleware"
	_movieHandler "github.com/minhthong582000/go-movies-api/movie/delivery/http"
	_gormMovieRepo "github.com/minhthong582000/go-movies-api/movie/repository/gorm"
	_movieUsecase "github.com/minhthong582000/go-movies-api/movie/usecase"

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
	db.AutoMigrate(&domain.Movie{}, &domain.Director{}, &domain.Actor{}, &domain.Gendres{}, &domain.TrailerLink{})

	// Init auth http middleware
	authUse := _authUsecase.NewAuthUsecase()
	authHTTPMdw := _authHttpMiddleware.NewAuthMiddleware(authUse)

	// Init movie handler
	movieRepo := _gormMovieRepo.NewGormMovieRepository(db)
	movieUse := _movieUsecase.NewMovieUsecase(movieRepo)
	r.Use(authHTTPMdw.AuthorizeJWT(), func(c *gin.Context) {
		_movieHandler.NewMovieHandler(r, movieUse)
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
