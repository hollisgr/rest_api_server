package user_repository

import "rest_api/internal/service/dto"

type UserRepository interface {
	AuthUser(userWeb dto.WebUserAuth) (token string, err error)
	CreateUser(userWeb dto.WebUserCreate) (id int, err error)
	LoadUserByID(id int) (userWeb dto.WebUserLoad, err error)
	LoadUserByLogin(login string) (userWeb dto.WebUserLoad, err error)
	LoadUserList() (usersWeb []dto.WebUserListLoad, err error)
	UpdateUser(userWeb dto.WebUserUpdate) (err error)
	DeleteUser(id int) (err error)
}
