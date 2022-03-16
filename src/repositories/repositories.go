package repositories

import (
	"github.com/jhamiltonjunior/ecommerce-backend/src/repositories/users"
	"github.com/jmoiron/sqlx"
)

type Container struct {
	User users.UserRepositories
}

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB

	// WriterGorm *gorm.DB ???
	// ReaderGorm *gorm.DB ???
}

func New(opts Options) *Container {
	return &Container{
		User: users.NewSqlxRespository(opts.WriterSqlx, opts.ReaderSqlx),
	}
}
