package repository

import (
	"context"

	"github.com/osmait/admin-finanzas/models"
)

type Repository interface {
	// Users
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByToken(ctx context.Context, token string) (*models.User, error)
	UpdateConfirmed(ctx context.Context, user *models.User) error

	// account
	InsertAccount(ctx context.Context, account *models.Account) error
	GetAcoounts(ctx context.Context, userId string) ([]*models.Account, error)
	DeleteAccount(ctx context.Context, id string, userId string) error
	// Bills

	// Income
	InsertIncome(ctx context.Context, transaction *models.Transaction) error

	GetIncome(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Transaction, error)
	GetAllTransaction(ctx context.Context, userId string, date1 string, date2 string) ([]*models.Transaction, error)
	UpdateTransaction(ctx context.Context, id string, transaction *models.Transaction) error
	GetBalace(ctx context.Context, id string) ([]*models.Balace, error)
	DeleteIncome(ctx context.Context, id string) error

	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

// Users

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}
func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}
func GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserEmail(ctx, email)
}
func GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	return implementation.GetUserByToken(ctx, token)
}
func UpdateConfirmed(ctx context.Context, user *models.User) error {
	return implementation.UpdateConfirmed(ctx, user)
}

func Close() error {
	return implementation.Close()
}

// Account

func InsertAccount(ctx context.Context, acount *models.Account) error {
	return implementation.InsertAccount(ctx, acount)
}

func GetAcoounts(ctx context.Context, userId string) ([]*models.Account, error) {
	return implementation.GetAcoounts(ctx, userId)
}

func DeleteAccount(ctx context.Context, id string, userId string) error {
	return implementation.DeleteAccount(ctx, id, userId)
}

// Income
func InsertIncome(ctx context.Context, transaction *models.Transaction) error {
	return implementation.InsertIncome(ctx, transaction)
}
func GetAllTransaction(ctx context.Context, userId string, date1 string, date2 string) ([]*models.Transaction, error) {
	return implementation.GetAllTransaction(ctx, userId, date1, date2)
}

func GetIncome(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Transaction, error) {
	return implementation.GetIncome(ctx, accountId, date1, date2)
}

func UpdateTransaction(ctx context.Context, id string, transaction *models.Transaction) error {
	return implementation.UpdateTransaction(ctx, id, transaction)
}

func DeleteIncome(ctx context.Context, id string) error {
	return implementation.DeleteIncome(ctx, id)
}

func GetBalance(ctx context.Context, id string) ([]*models.Balace, error) {
	return implementation.GetBalace(ctx, id)
}
