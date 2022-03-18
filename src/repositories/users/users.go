package users

import (
	"context"

	"github.com/jhamiltonjunior/ecommerce-backend/src/entities"
)

type UserRepositories interface {
	Create(ctx context.Context, newUser entities.User) (interface{}, error)
	GetById(ctx context.Context, ID int) (*entities.UserWithoutPassword, error)

	UpdateById(ctx context.Context, ID int, user entities.User) (error)
	DeleteById(ID int) error

	// GetAll(ctx context.Context) ([]entities.User, error)
}
