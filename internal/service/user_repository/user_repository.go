package user_repository

import "rest_api/internal/service/dto"

type UserRepository interface {
	CreateUser(u dto.UserCreateDTO) (id int, err error)
	LoadUserByID(id int) (u dto.UserLoadDTO, err error)
	LoadUserByLogin(login string) (u dto.UserLoadDTO, err error)
	LoadUserList() (userArray []dto.UserListLoadDTO, err error)
	UpdateUser(u dto.UserUpdateDTO) (err error)
	DeleteUser(id int) (err error)
}
