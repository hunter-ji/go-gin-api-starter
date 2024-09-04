// @Title user.go
// @Description
// @Author Hunter 2024/9/4 16:10

package service

import (
	"errors"

	"go-gin-api-starter/internal/database"
	"go-gin-api-starter/internal/model"
	"go-gin-api-starter/internal/repository"
	"go-gin-api-starter/pkg/auth"
	"go-gin-api-starter/pkg/util/crypto"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Login(req *model.LoginRequest) (*model.AuthResponse, error) {
	user, err := s.userRepo.GetUserByAccountName(req.AccountName)
	if err != nil {
		return nil, err
	}

	if user.Password != crypto.Md5(req.Password) {
		return nil, errors.New("invalid password")
	}

	accessToken, refreshToken, err := auth.GenerateAccessTokenAndRefreshToken(user.ID, database.RDB)
	if err != nil {
		return nil, err
	}

	resp := &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return resp, nil
}

func (s *UserService) Logout(userID uint64) error {
	err := auth.DeleteToken(userID, database.RDB)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) RefreshToken(refreshToken string) (string, string, error) {
	newAccessToken, newRefreshToken, err := auth.RefreshToken(refreshToken, database.RDB)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *UserService) GetUserByID(userID uint64) (*model.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
