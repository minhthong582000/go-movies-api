package domain

import (
	"context"
	"time"
)

// Director ...
type Director struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Bio       string    `json:"bio"`
	Image     string    `gorm:"type:varchar(256);UNIQUE" json:"url" binding:"required,url"`
	Awards    string    `json:"awards"`
	Movie     []*Movie  `gorm:"many2many:movie_directors;" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DirectorRepository represent the author's repository contract
type DirectorRepository interface {
	GetByID(ctx context.Context, id int64) (Director, error)
}
