package repositories

import (
	"context"
	"golang-postgresql/model"
)

type UserPostgresRepository interface {
	Insert(ctx context.Context, user model.User) error

	Update(ctx context.Context, user model.User) error

	Delete(ctx context.Context, user model.User) error

	FindAll(ctx context.Context) ([]model.User, error)
}
