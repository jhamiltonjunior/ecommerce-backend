package repositories

import (
	"github.com/jackc/pgx/v4"
	"github.com/jhamiltonjunior/ecommerce-backend/src/repositories/users"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	User users.UserRepositories
}

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
	
	WriterPgx *pgx.Conn
	ReaderPgx *pgx.Conn

	// WriterGorm *gorm.DB ???
	// ReaderGorm *gorm.DB ???
}

func New(opts Options) *Container {
	return &Container{
		User: users.NewPgxRepository(opts.WriterPgx, opts.ReaderPgx),
	}
}
