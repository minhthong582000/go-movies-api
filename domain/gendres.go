package domain

// Gendres ...
type Gendres struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name" binding:"required"`
	Movie []*Movie `gorm:"many2many:movie_gendres;" json:"-"`
}
