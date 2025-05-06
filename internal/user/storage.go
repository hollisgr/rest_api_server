package user

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) error
	FindUser(ctx context.Context, id int64) (User, error)
	FindAllUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, id int64) error
	FindAllTgUsers(tcx context.Context) ([]TGUser, error)
	CreateTgUser(ctx context.Context, user TGUser) error
	SetAdmin(ctx context.Context, tg_id int64) error
	FindTgUser(ctx context.Context, tg_id int64) (TGUser, error)
}
