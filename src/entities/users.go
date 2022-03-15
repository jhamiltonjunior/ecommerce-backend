package entities

import "time"

type User struct {
	ID uint `json:"user_id" gorm:"id; primaryKey; autoIncrement;"`

	// I put Name, because if I put UserName when going to use
	// would have to call user.UserName and I don't like that
	// user.Name is already implied
	//  used for get in user account
	// UserName string `json:"username" gorm:"unique; not null"`

	// gorm.Model

	// this is full name of people
	// not is used for get in
	FullName string `json:"full_name"`
	// also used for get in user account
	Email    string `json:"email" gorm:"email; type: varchar(100); unique; not null"`
	Password []byte `json:"password" gorm:"password; not null"`

	// the timestamp
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}