package user_service

import (
	"context"
	"rest_api/internal/db/storage"
	"rest_api/internal/service/dto"
	"rest_api/internal/service/mappers"
	"rest_api/internal/service/user_repository"
	"rest_api/internal/service/validate"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	storage   storage.Storage
	validator *validator.Validate
}

func NewUserService(s storage.Storage, v *validator.Validate) user_repository.UserRepository {
	return &UserService{
		storage:   s,
		validator: v,
	}
}

func (s *UserService) CreateUser(u dto.UserCreateDTO) (id int, err error) {
	err = s.validator.Struct(u)
	if err != nil {
		return id, err
	}

	err = validate.CheckPassword(u.Password)

	if err != nil {
		return id, err
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return id, err
	}

	userDB := mappers.FromCreateUserToDB(u)
	userDB.PasswordHash = string(pwdHash)

	id, err = s.storage.CreateUser(context.Background(), userDB)
	return id, err
}

func (s *UserService) LoadUserByID(id int) (u dto.UserLoadDTO, err error) {
	u, err = s.storage.LoadUserByID(context.Background(), id)
	return u, err
}

func (s *UserService) LoadUserByLogin(login string) (u dto.UserLoadDTO, err error) {
	u, err = s.storage.LoadUserByLogin(context.Background(), login)
	return u, err
}

func (s *UserService) UpdateUser(u dto.UserUpdateDTO) (err error) {
	err = s.validator.Struct(u)
	if err != nil {
		return err
	}

	err = s.storage.UpdateUser(context.Background(), u)
	return err
}

func (s *UserService) DeleteUser(id int) (err error) {
	err = s.storage.DeleteUser(context.Background(), id)
	return err
}

func (s *UserService) LoadUserList() (userArray []dto.UserListLoadDTO, err error) {
	userArray, err = s.storage.LoadAllUsers(context.Background())
	return userArray, err
}
