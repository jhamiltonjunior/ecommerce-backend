package configs

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

func GetWriterPgx() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		return nil
	}

	// defer conn.Close(context.Background())

	return conn
}

func GetReaderPgx() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		return nil
	}

	// defer conn.Close(context.Background())

	return conn
}
