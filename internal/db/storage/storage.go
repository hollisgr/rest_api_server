package storage

import (
	"context"
	"rest_api/internal/service/dto"
)

type Storage interface {
	CreateUser(ctx context.Context, u dto.DBUserCreateDTO) (int, error)
	LoadUserByID(ctx context.Context, id int) (dto.UserLoadDTO, error)
	LoadUserByLogin(ctx context.Context, login string) (dto.UserLoadDTO, error)
	LoadAllUsers(ctx context.Context) ([]dto.UserListLoadDTO, error)
	UpdateUser(ctx context.Context, u dto.UserUpdateDTO) error
	DeleteUser(ctx context.Context, id int) error
}
