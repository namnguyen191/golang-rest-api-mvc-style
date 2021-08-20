package service

import (
	"log"

	"github.com/namnguyen191/themuzix-golang-rest-api/dto"
	"github.com/namnguyen191/themuzix-golang-rest-api/entity"
	"github.com/namnguyen191/themuzix-golang-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func comparePassword(hashedPwd []byte, plainPwd []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPwd, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword([]byte(v.Password), []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	return service.userRepository.InsertUser(userToCreate)
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}
