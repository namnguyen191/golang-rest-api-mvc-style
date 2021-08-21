package service

import (
	"github.com/namnguyen191/themuzix-golang-rest-api/dto"
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"github.com/namnguyen191/themuzix-golang-rest-api/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	updatedUser := s.userRepo.UpdateUser(userToUpdate)

	return updatedUser
}

func (s *userService) Profile(userID string) entity.User {
	return s.userRepo.ProfileUser(userID)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}
