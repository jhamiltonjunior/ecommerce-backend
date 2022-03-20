package users

import (
	"context"

	"github.com/jhamiltonjunior/ecommerce-backend/src/entities"
)

type UserRepositories interface {
	Create(ctx context.Context, newUser entities.User) (interface{}, error)
	GetById(ctx context.Context, ID int64) (*entities.UserWithoutPassword, error)

	UpdateById(ctx context.Context, ID int64, user *entities.User) (*entities.User, error)
	DeleteById(ID int64) error

	// GetAll(ctx context.Context) ([]entities.User, error)
}
