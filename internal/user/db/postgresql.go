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
	id := 0
	r.client.QueryRow(ctx, countQuery).Scan(&count)
	user.ID = int64(count) + 1

	PasswordHash, cryptErr := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if cryptErr != nil {
		return fmt.Errorf("failed to create user hash to error: %v", cryptErr)
	}

	user.PasswordHash = string(PasswordHash)

	query := `
			INSERT INTO users (id, username, email, passwordhash) RETURNING id
			VALUES ($1, $2, $3, $4)
			`
	r.logger.Traceln("SQL Query:", formatQuery(query))
	r.client.QueryRow(ctx, query, user.ID, user.Username, user.Email, user.PasswordHash, &id)
	if id == 0 {
		return fmt.Errorf("failed to create user, args not unique")
	}

	return nil
}
