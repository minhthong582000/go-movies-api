package usecase

import (
	"github.com/minhthong582000/go-movies-api/domain"
)

type movieUsecase struct {
	movieRepo domain.MovieRepository
}

func NewMovieUsecase(mv domain.MovieRepository) domain.MovieUsecase {
	return &movieUsecase{
		movieRepo: mv,
	}
}

func (m *movieUsecase) Fetch() (movies []domain.Movie, err error) {
	movies, err = m.movieRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (m *movieUsecase) GetByID(id int64) (movie domain.Movie, err error) {
	movie, err = m.movieRepo.GetByID(id)
	if err != nil {
		return domain.Movie{}, err
	}

	return movie, nil
}

func (m *movieUsecase) Update(id int64, movie *domain.Movie) (err error) {
	err = m.movieRepo.Update(id, movie)
	if err != nil {
		return err
	}

	return nil
}

func (m *movieUsecase) GetByTitle(title string) (movie domain.Movie, err error) {
	movie, err = m.movieRepo.GetByTitle(title)
	if err != nil {
		return domain.Movie{}, err
	}

	return movie, err
}

func (m *movieUsecase) Store(movie *domain.Movie) (err error) {
	existed, _ := m.movieRepo.GetByTitle(movie.Title)
	if existed == (domain.Movie{}) {
		return domain.ErrConflict
	}

	return m.movieRepo.Store(movie)
}

func (m *movieUsecase) Delete(id int64) error {
	existed, _ := m.movieRepo.GetByID(id)
	if existed == (domain.Movie{}) {
		return domain.ErrNotFound
	}

	return m.movieRepo.Delete(id)
}
