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

func (repo *repoSqlx) GetById(ctx context.Context, ID int) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			full_name,
			email
		FROM users
		WHERE id=$1
	`, ID)
	// created_at,
	// updated_at
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *repoSqlx) UpdateById(ctx context.Context, ID int, user entities.User) error {

	_, err := repo.reader.ExecContext(ctx, `
		UPDATE users
		SET
			full_name = $2,
			email = $3,
			password = $4,
			updated_at = $5
		WHERE id=$1
	`, ID, user.FullName, user.Email, user.Password, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
