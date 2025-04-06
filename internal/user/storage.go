package user

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) (string, error)
	FindUser(ctx context.Context, id string) (User, error)
	DeleteUser(ctx context.Context, id string) error
}
