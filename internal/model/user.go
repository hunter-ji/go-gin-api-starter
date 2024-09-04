// @Title user.go
// @Description
// @Author Hunter 2024/9/4 15:57

package model

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type User struct {
	ID          uint64                `json:"id" gorm:"column:id;primary_id"`
	AccountName string                `json:"account_name" gorm:"column:account_name"`
	Password    string                `json:"password" gorm:"column:password"`
	CreatedAt   time.Time             `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time             `json:"updated_at" gorm:"column:updated_at"`
	IsDeleted   soft_delete.DeletedAt `json:"-" gorm:"column:is_deleted;softDelete:flag"`
}

type LoginRequest struct {
	AccountName string `json:"account_name" binding:"required,max=255"`
	Password    string `json:"password" binding:"required,max=255"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required,max=1000"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type GetUserInfoRequest struct {
	ID uint64 `uri:"id" binding:"required,min=1"`
}
