package util

import (

	"github.com/jhamiltonjunior/ecommerce-backend/src/configs"
	"github.com/jhamiltonjunior/ecommerce-backend/src/repositories"
)

func NewRepositories() *repositories.Container {
	repos := repositories.New(repositories.Options{
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterPgx:  configs.GetWriterPgx(),
		ReaderPgx:  configs.GetReaderPgx(),
	})

	return repos
}