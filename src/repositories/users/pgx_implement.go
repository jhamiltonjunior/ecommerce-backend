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
	row, err := repo.writer.Query(ctx, `
		INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING *
	`, newUser.FullName, newUser.Email, newUser.Password)
	if err != nil {
		return 0, err
	}
	defer row.Close()
	row.Next()
	// fmt.Println(row.Values())
	
	var id interface{}
	// row.Next()
	items, _ := row.Values()
	id = items[0]

	fmt.Println(items[0])

	return id, nil
}

func (repo *repoPgx) GetById(ctx context.Context, ID int) (*entities.UserWithoutPassword, error) {
	var user *entities.UserWithoutPassword

	return user, nil
}

func (repo *repoPgx) UpdateById(ctx context.Context, ID int, user entities.User) error {
	panic("This not inplemented!")
}
func (repo *repoPgx) DeleteById(ID int) error {
	panic("This not inplemented!")
}
