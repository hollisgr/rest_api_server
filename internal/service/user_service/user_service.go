package user_service

import (
	"context"
	"fmt"
	"rest_api/internal/db/storage"
	"rest_api/internal/entities/user"
	"rest_api/internal/service/dto"
	"rest_api/internal/service/jwt"
	"rest_api/internal/service/mappers"
	"rest_api/internal/service/user_repository"
	"rest_api/internal/service/validate"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

const (
	FAIL = 0
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

func (s *UserService) CreateUser(userWeb dto.WebUserCreate) (id int, err error) {
	err = s.validator.Struct(userWeb)
	if err != nil {
		return FAIL, err
	}

	err = validate.CheckPassword(userWeb.Password)
	if err != nil {
		return FAIL, err
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(userWeb.Password), 10)
	if err != nil {
		return FAIL, err
	}

	user := mappers.FromWebUserCreateToUser(userWeb)
	user.PasswordHash = string(pwdHash)

	userDB := mappers.FromUserToDBUserCreate(user)

	err = s.validator.Struct(userDB)
	if err != nil {
		return FAIL, err
	}

	id, err = s.storage.CreateUser(context.Background(), userDB)

	if err != nil {
		return FAIL, err
	}

	return id, nil
}

func (s *UserService) LoadUserByID(id int) (userWeb dto.WebUserLoad, err error) {

	userDB, err := s.storage.LoadUserByID(context.Background(), id)

	if err != nil {
		return userWeb, err
	}

	err = s.validator.Struct(userDB)
	if err != nil {
		return userWeb, err
	}

	user := mappers.FromDBUserLoadToUser(userDB)

	userWeb = mappers.FromUserToWebUserLoad(user)

	err = s.validator.Struct(userWeb)
	if err != nil {
		return userWeb, err
	}

	return userWeb, nil
}

func (s *UserService) LoadUserByLogin(login string) (userWeb dto.WebUserLoad, err error) {
	userDB, err := s.storage.LoadUserByLogin(context.Background(), login)

	if err != nil {
		return userWeb, err
	}

	err = s.validator.Struct(userDB)
	if err != nil {
		return userWeb, err
	}

	user := mappers.FromDBUserLoadToUser(userDB)

	userWeb = mappers.FromUserToWebUserLoad(user)

	err = s.validator.Struct(userWeb)
	if err != nil {
		return userWeb, err
	}

	return userWeb, nil
}

func (s *UserService) UpdateUser(userWeb dto.WebUserUpdate) (err error) {
	err = s.validator.Struct(userWeb)
	if err != nil {
		return fmt.Errorf("%v, web user val err", err)
	}

	user := mappers.FromWebUserUpdateToUser(userWeb)

	userDB := mappers.FromUserToDBUserUpdate(user)

	err = s.validator.Struct(userDB)
	if err != nil {
		return fmt.Errorf("%v, db user val err", err)
	}

	err = s.storage.UpdateUser(context.Background(), userDB)

	if err != nil {
		return fmt.Errorf("%v, storage update user err", err)
	}

	return nil
}

func (s *UserService) DeleteUser(id int) (err error) {
	err = s.storage.DeleteUser(context.Background(), id)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) LoadUserList() (usersWeb []dto.WebUserListLoad, err error) {
	usersDB, err := s.storage.LoadAllUsers(context.Background())

	if err != nil {
		return nil, err
	}

	for _, userDB := range usersDB {
		err = s.validator.Struct(userDB)
		if err != nil {
			return nil, err
		}
	}

	users := make([]user.User, 0)

	for _, userDB := range usersDB {
		tempUser := mappers.FromDBUserListLoadToUser(userDB)
		users = append(users, tempUser)
	}

	usersWeb = make([]dto.WebUserListLoad, 0)

	for _, user := range users {
		tempUserWeb := mappers.FromUserToWebUserListLoad(user)
		usersWeb = append(usersWeb, tempUserWeb)
	}

	return usersWeb, nil
}

func (s *UserService) AuthUser(userWeb dto.WebUserAuth) (token string, err error) {
	userDB, err := s.storage.LoadUserByLogin(context.Background(), userWeb.Login)
	if err != nil {
		return token, err
	}

	err = s.validator.Struct(userDB)
	if err != nil {
		return token, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.PasswordHash), []byte(userWeb.Password))
	if err != nil {
		return token, err
	}

	user := mappers.FromDBUserLoadToUser(userDB)

	jwtCreate := mappers.FromUserToJWTTokenCreate(user)

	token, err = jwt.CreateToken(jwtCreate)
	if err != nil {
		return token, err
	}

	return token, nil
}
