package repository

import (
	"context"

	"github.com/osmait/admin-finanzas/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserEmail(ctx context.Context, email string) (*models.User, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
func GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserEmail(ctx, email)
}

func Close() error {
	return implementation.Close()
}
