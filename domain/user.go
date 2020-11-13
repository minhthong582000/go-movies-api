package domain

import (
	"time"
)

// User model
type User struct {
	ID        int64     `json:"-" gorm:"primary_key;auto_increment"`
	Username  string    `json:"username" gorm:"type:varchar(40);unique"`
	Password  string    `json:"password" gorm:"type:varchar(100)"`
	FullName  string    `json:"fullName" gorm:"type:varchar(100)"`
	Age       uint16    `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	GetByID(id int64) (User, error)
	GetByUsername(username string) (User, error)
	Store(user *User) error
	Update(id int64, user *User) (updatedRow int64, err error)
	Delete(id int64) error
	ComparePassword(password string, hashedPassword string) (isValid bool, err error)
}
