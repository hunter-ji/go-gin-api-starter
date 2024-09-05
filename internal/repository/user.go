// @Title user.go
// @Description
// @Author Hunter 2024/9/4 16:04

package repository

import (
	"go-gin-api-starter/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByAccountName(accountName string) (*model.User, error) {
	var user model.User
	result := r.db.Where("account_name = ?", accountName).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(userID uint64) (*model.User, error) {
	var user model.User
	result := r.db.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
