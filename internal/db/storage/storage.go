package storage

import (
	"context"
	"rest_api/internal/service/dto"
)

type Storage interface {
	CreateUser(ctx context.Context, u dto.DBUserCreate) (int, error)
	LoadUserByID(ctx context.Context, id int) (dto.DBUserLoad, error)
	LoadUserByLogin(ctx context.Context, login string) (dto.DBUserLoad, error)
	LoadAllUsers(ctx context.Context) ([]dto.DBUserListLoad, error)
	UpdateUser(ctx context.Context, u dto.DBUserUpdate) error
	DeleteUser(ctx context.Context, id int) error
}
