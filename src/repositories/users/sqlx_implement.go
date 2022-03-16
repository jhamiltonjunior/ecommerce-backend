package users

import (
	"context"

	"github.com/jhamiltonjunior/ecommerce-backend/src/entities"
	"github.com/jmoiron/sqlx"
)

type repoSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRespository(w *sqlx.DB, r *sqlx.DB) UserRepositories {
	return &repoSqlx{writer: w, reader: r}
}

func (repo *repoSqlx) Create(ctx context.Context, newUser entities.User) error {
	_, err := repo.writer.ExecContext(ctx, `
		INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING *
	`, newUser.FullName, newUser.Email, newUser.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repoSqlx) GetById(ctx context.Context, ID int64) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			full_name,
			email
		FROM users
		WHERE id=$1
	`, ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}