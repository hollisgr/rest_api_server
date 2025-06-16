package mappers

import (
	"rest_api/internal/entities/user"
	"rest_api/internal/service/dto"
)

// USER TO DB

func FromUserToDBUserCreate(u user.User) dto.DBUserCreate {
	res := dto.DBUserCreate{
		Login:        u.Login,
		FirstName:    u.FirstName,
		SecondName:   u.SecondName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
	return res
}

func FromUserToDBUserUpdate(u user.User) dto.DBUserUpdate {
	res := dto.DBUserUpdate{
		Id:         u.Id,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}

// USER TO WEB

func FromUserToWebUserLoad(u user.User) dto.WebUserLoad {
	res := dto.WebUserLoad{
		Id:         u.Id,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}

func FromUserToWebUserListLoad(u user.User) dto.WebUserListLoad {
	res := dto.WebUserListLoad{
		Id:         u.Id,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}

// WEB USER TO USER

func FromWebUserCreateToUser(u dto.WebUserCreate) user.User {
	res := user.User{
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}

func FromWebUserUpdateToUser(u dto.WebUserUpdate) user.User {
	res := user.User{
		Id:         u.Id,
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}

// DBUSER TO USER

func FromDBUserLoadToUser(u dto.DBUserLoad) user.User {
	res := user.User{
		Id:           u.Id,
		Login:        u.Login,
		FirstName:    u.FirstName,
		SecondName:   u.SecondName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
	return res
}

func FromDBUserListLoadToUser(u dto.DBUserListLoad) user.User {
	res := user.User{
		Id:           u.Id,
		Login:        u.Login,
		FirstName:    u.FirstName,
		SecondName:   u.SecondName,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
	return res
}

// JWT TOKEN CREATE

func FromUserToJWTTokenCreate(u user.User) dto.JWTTokenCreate {
	res := dto.JWTTokenCreate{
		Login:      u.Login,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
	}
	return res
}
