package gorm

import (
	"github.com/minhthong582000/go-movies-api/domain"
	"gorm.io/gorm"
)

type gormMovieRepository struct {
	Conn *gorm.DB
}

// NewGormMovieRepository will create an object that represent the movie.Repository interface
func NewGormMovieRepository(Conn *gorm.DB) domain.MovieRepository {
	return &gormMovieRepository{Conn}
}

func (m *gormMovieRepository) Fetch() ([]domain.Movie, error) {
	movies := make([]domain.Movie, 0)
	if err := m.Conn.Joins("Author").Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *gormMovieRepository) GetByID(id int64) (domain.Movie, error) {
	var movie domain.Movie
	if err := m.Conn.Joins("Author").Where("movies.id = ?", id).Find(&movie).Error; err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}
func (m *gormMovieRepository) GetByTitle(title string) (res domain.Movie, err error) {
	var movie domain.Movie
	if err := m.Conn.Find(&movie).Where("title = ?", title).Error; err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}

func (m *gormMovieRepository) Store(movie *domain.Movie) error {
	if err := m.Conn.Create(&movie).Error; err != nil {
		return err
	}

	return nil
}

func (m *gormMovieRepository) Delete(id int64) (err error) {
	var movie domain.Movie
	if err := m.Conn.Delete(&movie).Where("id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (m *gormMovieRepository) Update(id int64, movie *domain.Movie) error {
	if err := m.Conn.Model(domain.Movie{}).Where("id = ?", id).Updates(&movie).Error; err != nil {
		return err
	}

	return nil
}
