package domain

import (
	"time"
)

// Movie struct
type Movie struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(100)" json:"title" binding:"min=2,max=10"`
	Description string    `gorm:"type:varchar(100)" json:"description" binding:"max=20"`
	URL         string    `gorm:"type:varchar(100);UNIQUE" json:"url" binding:"required,url"`
	Author      Author    `gorm:"foreignkey:AuthorID" json:"author" binding:"required"`
	AuthorID    int64     `json:"authorID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MovieUsecase interface
type MovieUsecase interface {
	Fetch() (movies []Movie, err error)
	GetByID(id int64) (movie Movie, err error)
	Update(id int64, movie *Movie) (err error)
	Store(movie *Movie) (err error)
	Delete(id int64) error
}

// MovieRepository interface
type MovieRepository interface {
	Fetch() (res []Movie, err error)
	GetByID(id int64) (Movie, error)
	Update(id int64, movie *Movie) error
	GetByTitle(title string) (Movie, error)
	Store(movie *Movie) error
	Delete(id int64) error
}
