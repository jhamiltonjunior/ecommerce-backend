package util

import (
	"fmt"

	"github.com/jhamiltonjunior/ecommerce-backend/src/configs"
	"github.com/jhamiltonjunior/ecommerce-backend/src/repositories"
)

// var (
// 	errDuplicateEmail = fmt.Errorf(
// 		`ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)`,
// 	)
// )

func NewRepositories() *repositories.Container {
	repos := repositories.New(repositories.Options{
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterPgx:  configs.GetWriterPgx(),
		ReaderPgx:  configs.GetReaderPgx(),
	})

	return repos
}

func IsDuplicate(value string) string {
	return fmt.Sprintf(
		`ERROR: duplicate key value violates unique constraint "%v" (SQLSTATE 23505)`, value,
	)

	// return errDuplicateValue
}

// duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)