package configs

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

func GetReaderPgx() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		return nil
	}

	// defer conn.Close(context.Background())

	return conn
}

// func GetWriterPgx() *sqlx.DB {
// 	pgx
// 	writer := sqlx.MustConnect("postgres", os.Getenv("DATA_SOURCE_NAME"))
// 	return writer
// }
