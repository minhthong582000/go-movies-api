package gorm

import (
	"github.com/minhthong582000/go-movies-api/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	Conn *gorm.DB
}

func NewGormUserRepository(conn *gorm.DB) domain.UserRepository {
	return &gormUserRepository{
		Conn: conn,
	}
}

func (u gormUserRepository) GetByID(id int64) (domain.User, error) {
	var user domain.User
	if err := u.Conn.Where("id = ?", id).Find(&user).Error; err != nil {
		return domain.User{}, nil
	}

	return user, nil
}

func (u gormUserRepository) GetByUsername(username string) (domain.User, error) {
	var user domain.User
	if err := u.Conn.Where("username = ?", username).Find(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (u gormUserRepository) Store(user *domain.User) error {
	if err := u.Conn.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u gormUserRepository) Update(id int64, user *domain.User) (updatedRow int64, err error) {
	result := u.Conn.Where("id = ?", id).Updates(&user)

	if err := result.Error; err != nil {
		return 0, err
	}

	return result.RowsAffected, nil
}

func (u gormUserRepository) Delete(id int64) error {
	var user domain.User
	if err := u.Conn.Delete(&user).Where("id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (u gormUserRepository) ComparePassword(password string, hashedPassword string) (isValid bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
