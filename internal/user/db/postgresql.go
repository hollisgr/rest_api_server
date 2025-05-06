package db

import (
	"context"
	"fmt"
	"rest_api_server/internal/user"
	"rest_api_server/pkg/client/postgresql"
	"rest_api_server/pkg/logging"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) FindUser(ctx context.Context, id int64) (user.User, error) {
	user := user.User{}

	query := ` 
		SELECT 
			id, username, email
		FROM users 
		WHERE id = $1
		`

	r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindAllUsers(ctx context.Context) ([]user.User, error) {
	query := `
		SELECT 
			id, username, email
		FROM users
		`
	users := make([]user.User, 0)

	r.logger.Traceln("SQL Query:", formatQuery(query))
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tempUser := user.User{}
		err = rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUser)
	}
	return users, nil
}

func (r *repository) DeleteUser(ctx context.Context, id int64) error {
	ID := 0
	query := `
		DELETE FROM users 
		WHERE id = $1
		RETURNING id
		`
	r.client.QueryRow(ctx, query, id).Scan(&ID)
	r.logger.Traceln("SQL Query:", formatQuery(query))

	if ID == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *repository) CreateUser(ctx context.Context, user user.User) error {

	r.logger.Infoln("creating user")
	countQuery := `
		SELECT
			MAX(ID)
		FROM users
		`
	count := 0
	r.client.QueryRow(ctx, countQuery).Scan(&count)
	user.ID = int64(count) + 1

	PasswordHash, cryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if cryptErr != nil {
		return fmt.Errorf("failed to create user hash to error: %v", cryptErr)
	}

	user.PasswordHash = string(PasswordHash)
	var id int
	query := `
			INSERT INTO users (id, username, email, passwordhash)
			VALUES ($1, $2, $3, $4)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, user.ID, user.Username, user.Email, user.PasswordHash).Scan(&id)
	if id == 0 {
		return fmt.Errorf("failed to create user, args not unique")
	}

	return nil
}

func (r *repository) CreateTgUser(ctx context.Context, user user.TGUser) error {
	r.logger.Infoln("creating tg_user")
	var id int
	default_role_id := 1
	query := `
			INSERT INTO tg_users (username, tg_id, role_id)
			VALUES ($1, $2, $3)
			RETURNING id
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, user.TG_USERNAME, user.TG_ID, default_role_id).Scan(&id)
	if id == 0 {
		return fmt.Errorf("failed to create user, args not unique")
	}

	return nil
}

func (r *repository) FindTgUser(ctx context.Context, tg_id int64) (user.TGUser, error) {
	user := user.TGUser{}

	query := ` 
		SELECT 
			id, username, tg_id, role_id
		FROM tg_users 
		WHERE tg_id = $1
		`

	r.logger.Traceln("SQL Query:", formatQuery(query))
	row := r.client.QueryRow(ctx, query, tg_id)
	err := row.Scan(&user.ID, &user.TG_USERNAME, &user.TG_ID, &user.Role_id)
	if err != nil {
		return user, err
	}

	fmt.Println(user)

	return user, nil
}

func (r *repository) FindAllTgUsers(ctx context.Context) ([]user.TGUser, error) {
	query := `
		SELECT 
			id, username, tg_id, role_id
		FROM tg_users
		`
	users := make([]user.TGUser, 0)

	r.logger.Traceln("SQL Query:", formatQuery(query))
	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		tempUser := user.TGUser{}
		err = rows.Scan(&tempUser.TG_ID, &tempUser.TG_USERNAME, &tempUser.Role_id)
		if err != nil {
			return nil, err
		}
		users = append(users, tempUser)
	}
	return users, nil
}

func (r *repository) SetAdmin(ctx context.Context, tg_id int64) error {
	ID := 0
	query := `
		UPDATE 
			tg_users 
		SET 
			role_id = 2
		WHERE 
			tg_id = $1
		RETURNING id
		`
	r.client.QueryRow(ctx, query, tg_id).Scan(&ID)
	r.logger.Traceln("SQL Query:", formatQuery(query))

	if ID == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
