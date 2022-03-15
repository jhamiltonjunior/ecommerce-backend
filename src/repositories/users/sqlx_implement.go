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

func NewSqlxRespository(w *sqlx.DB, r *sqlx.DB) UserRepository {
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
