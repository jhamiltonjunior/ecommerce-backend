package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jhamiltonjunior/ecommerce-backend/src/entities"
)

type repoPgx struct {
	writer *pgx.Conn
	reader *pgx.Conn
}

func NewPgxRepository(w *pgx.Conn, r *pgx.Conn) UserRepositories {
	return &repoPgx{
		writer: w,
		reader: r,
	}
}

func (repo *repoPgx) Create(ctx context.Context, newUser entities.User) (interface{}, error) {
	// Estou usando query aqui porque eu preciso retornar o id par usar no JWT
	// isso permanecerá assim até que eu tenha outra solução
	rows, err := repo.writer.Query(ctx, `
		INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING *
	`, newUser.FullName, newUser.Email, newUser.Password)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id interface{}

	rows.Next()
	items, _ := rows.Values()
	id = items[0]

	fmt.Println(items[0])

	return id, nil
}

func (repo *repoPgx) GetById(ctx context.Context, id int64) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword

	err := repo.reader.QueryRow(ctx, `
		SELECT 
			id,
			full_name,
			email,
			created_at,
			updated_at
		FROM users
		WHERE id=$1
	`, id).Scan(
		&user.ID, &user.FullName,
		&user.Email, &user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *repoPgx) UpdateById(ctx context.Context, id int64, user *entities.User) (*entities.User, error) {
	rows, err := repo.writer.Query(ctx, `
		UPDATE users
		SET
			full_name = $2,
			email = $3,
			password = $4,
			updated_at = $5
		WHERE id=$1
		RETURNING *
	`, id, user.FullName, user.Email, user.Password, user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user.ID = id

	return user, nil
}
func (repo *repoPgx) DeleteById(ID int64) error {
	panic("This not inplemented!")
}
