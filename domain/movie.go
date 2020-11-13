package domain

import (
	"time"
)

// Movie struct
type Movie struct {
	ID           int64         `gorm:"primary_key;auto_increment" json:"id"`
	Title        string        `gorm:"type:varchar(100)" json:"title" binding:"min=2,max=10,required"`
	Description  string        `gorm:"type:varchar(100)" json:"description" binding:"max=20"`
	Year         int           `json:"year" binding:"min=1000"`
	Language     string        `gorm:"type:varchar(20)" json:"language" binding:"min=2,max=10"`
	IMDBScore    float32       `json:"IMDBScore" binding:"min=0,max=10"`
	TrailerLinks []TrailerLink `gorm:"foreignKey:MovieID" json:"trailerLinks,omitempty"`
	Directors    []*Director   `gorm:"many2many:movie_directors;" json:"directors,omitempty"`
	Actors       []*Actor      `gorm:"many2many:movie_actors;" json:"actors,omitempty"`
	Gendres      []*Gendres    `gorm:"many2many:movie_gendres;" json:"gendres,omitempty"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

// TrailerLink ...
type TrailerLink struct {
	ID      int64  `json:"id"`
	Link    string `json:"link"`
	Site    string `json:"site"`
	MovieID int64  `json:"movieID"`
}

// MovieUsecase interface
type MovieUsecase interface {
	Fetch() (movies []Movie, err error)
	GetByID(id int64) (movie Movie, err error)
	Update(id int64, movie *Movie) (updatedRow int64, err error)
	Store(movie *Movie) (err error)
	Delete(id int64) error
}

// MovieRepository interface
type MovieRepository interface {
	Fetch() (res []Movie, err error)
	GetByID(id int64) (Movie, error)
	Update(id int64, movie *Movie) (updatedRow int64, err error)
	GetByTitle(title string) (Movie, error)
	Store(movie *Movie) error
	Delete(id int64) error
}
