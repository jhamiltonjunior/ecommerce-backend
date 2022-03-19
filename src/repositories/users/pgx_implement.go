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

func (repo *repoPgx) GetById(ctx context.Context, id int) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword

	err := repo.reader.QueryRow(context.Background(), `
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

func (repo *repoPgx) UpdateById(ctx context.Context, ID int, user entities.User) error {
	panic("This not inplemented!")
}
func (repo *repoPgx) DeleteById(ID int) error {
	panic("This not inplemented!")
}
