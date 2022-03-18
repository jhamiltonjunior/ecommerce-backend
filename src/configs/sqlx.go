package configs

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetReaderSqlx() *sqlx.DB {
	reader := sqlx.MustConnect("postgres", os.Getenv("DATA_SOURCE_NAME"))

	return reader
}

func GetWriterSqlx() *sqlx.DB {
	writer := sqlx.MustConnect("postgres", os.Getenv("DATA_SOURCE_NAME"))
	return writer
}
