package db

import (
	"context"
	"fmt"
	"rest_api/internal/db/postgres"
	"rest_api/internal/db/storage"
	"rest_api/internal/service/dto"
)

type repository struct {
	client postgres.Client
}

// func formatQuery(q string) string {
// 	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
// }

func NewStorage(client postgres.Client) storage.Storage {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateUser(ctx context.Context, u dto.DBUserCreateDTO) (int, error) {
	id := 0

	query := `
		INSERT INTO 
			rest_api_users (
				login,
				first_name,
				second_name,
				email,
				password_hash
			)
		VALUES 
			($1, $2, $3, $4, $5)
		RETURNING 
			id
	`
	// log.Println("SQL Query:", formatQuery(query))

	r.client.QueryRow(ctx, query, u.Login, u.FirstName, u.SecondName, u.Email, u.PasswordHash).Scan(&id)

	if id == 0 {
		return id, fmt.Errorf("insert into table error")
	}
	return id, nil
}

func (r *repository) LoadUserByID(ctx context.Context, id int) (dto.UserLoadDTO, error) {
	query := `
		SELECT 
			id,
			login,
			first_name,
			second_name,
			email,
			password_hash
		FROM 
			rest_api_users
		WHERE 
			id = $1
	`

	// log.Println("SQL Query:", formatQuery(query))

	row := r.client.QueryRow(ctx, query, id)

	u := dto.UserLoadDTO{}
	err := row.Scan(&u.Id, &u.Login, &u.FirstName, &u.SecondName, &u.Email, &u.PasswordHash)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) LoadUserByLogin(ctx context.Context, login string) (dto.UserLoadDTO, error) {
	query := `
		SELECT 
			id,
			login,
			first_name,
			second_name,
			email,
			password_hash
		FROM 
			rest_api_users
		WHERE 
			login = $1
	`

	// log.Println("SQL Query:", formatQuery(query))

	row := r.client.QueryRow(ctx, query, login)

	u := dto.UserLoadDTO{}
	err := row.Scan(&u.Id, &u.Login, &u.FirstName, &u.SecondName, &u.Email, &u.PasswordHash)

	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) LoadAllUsers(ctx context.Context) ([]dto.UserListLoadDTO, error) {
	query := `
		SELECT 
			id,
			login,
			first_name,
			second_name,
			email
		FROM 
			rest_api_users
	`

	rows, err := r.client.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	dtoArr := make([]dto.UserListLoadDTO, 0)

	for rows.Next() {
		tempUser := dto.UserListLoadDTO{}

		err = rows.Scan(&tempUser.Id, &tempUser.Login, &tempUser.FirstName, &tempUser.SecondName, &tempUser.Email)
		if err != nil {
			return nil, err
		}

		dtoArr = append(dtoArr, tempUser)
	}

	return dtoArr, nil

}

func (r *repository) UpdateUser(ctx context.Context, u dto.UserUpdateDTO) error {
	id := u.Id
	query := `
		UPDATE 
			rest_api_users
		SET	
			login = $2,
			first_name = $3,
			second_name = $4,
			email = $5
		WHERE
			id = $1
		RETURNING id
	`
	row := r.client.QueryRow(ctx, query, id, u.Login, u.FirstName, u.SecondName, u.Email)
	err := row.Scan(&id)
	if err != nil || id != u.Id {
		return err
	}

	return nil
}

func (r *repository) DeleteUser(ctx context.Context, id int) error {
	query := `
		DELETE FROM
			rest_api_users
		WHERE
			id = $1
		RETURNING id
	`
	row := r.client.QueryRow(ctx, query, id)
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
