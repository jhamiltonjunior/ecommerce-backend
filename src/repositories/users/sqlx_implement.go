package users

import (
	"context"
	"fmt"

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

func (repo *repoSqlx) Create(ctx context.Context, newUser entities.User) (interface{}, error) {
	result, err := repo.writer.ExecContext(ctx, `
		INSERT INTO users (full_name, email, password) VALUES ($1, $2, $3) RETURNING *
	`, newUser.FullName, newUser.Email, newUser.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	fmt.Println(id)
	fmt.Println(err)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *repoSqlx) GetById(ctx context.Context, id int64) (*entities.UserWithoutPassword, error) {
	var user entities.UserWithoutPassword

	err := repo.reader.Get(&user, `
		SELECT 
			id,
			full_name,
			email
		FROM users
		WHERE id=$1
	`, id)
	// created_at,
	// updated_at
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *repoSqlx) UpdateById(ctx context.Context, id int64, user *entities.User) (*entities.User, error) {

	_, err := repo.writer.ExecContext(ctx, `
		UPDATE users
		SET
			full_name = $2,
			email = $3,
			password = $4,
			updated_at = $5
		WHERE id=$1
	`, id, user.FullName, user.Email, user.Password, user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// lastId, _ := result.LastInsertId()

	return nil, nil
}

func (repo *repoSqlx) DeleteById(id int64) error {
	if _, err := repo.writer.Exec("DELETE FROM users WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}
