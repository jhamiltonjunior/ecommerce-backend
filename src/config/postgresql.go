package config

// import (
// 	"database/sql"
// 	"os"

// 	"github.com/jmoiron/sqlx"
// 	_ "github.com/lib/pq"
// )

// func Open(dbsourse string) (db *sqlx.DB, err error) {
// 	db, err = sqlx.Open("postgres", dbsourse)
// 	if err != nil {
// 		return
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func Select(sql string) (*sql.Rows, error) {

// 	db, err := Open(
// 		os.Getenv("DB_SOURCE"),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := db.Query(sql)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// func Delete(sql string) (sql.Result, error) {
// 	db, err := Open(
// 		os.Getenv("DB_SOURCE"),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := db.Exec(sql)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }
