package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minhthong582000/go-movies-api/config"
	"github.com/minhthong582000/go-movies-api/domain"
	"github.com/minhthong582000/go-movies-api/middleware"
	movieHandler "github.com/minhthong582000/go-movies-api/movie/delivery/http"
	gormMovieRepo "github.com/minhthong582000/go-movies-api/movie/repository/gorm"
	movieUsecase "github.com/minhthong582000/go-movies-api/movie/usecase"

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

	movieRepo := gormMovieRepo.NewGormMovieRepository(db)
	movieUse := movieUsecase.NewMovieUsecase(movieRepo)
	movieHandler.NewMovieHandler(r, movieUse)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
