package controllers

import "fmt"

// var (
// 	errDuplicateEmail = fmt.Errorf(
// 		`ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)`,
// 	)
// )

func IsDuplicate(value string) string {
	return fmt.Sprintf(
		`ERROR: duplicate key value violates unique constraint "%v" (SQLSTATE 23505)`, value,
	)

	// return errDuplicateValue
}

// duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)