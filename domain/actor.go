package domain

import (
	"context"
	"time"
)

// Actor ...
type Actor struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Bio       string    `json:"bio"`
	Image     string    `gorm:"type:varchar(256);UNIQUE" json:"url" binding:"required,url"`
	Awards    string    `json:"awards"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ActorRepository represent the author's repository contract
type ActorRepository interface {
	GetByID(ctx context.Context, id int64) (Director, error)
}
