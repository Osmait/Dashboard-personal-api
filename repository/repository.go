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
	InsertBill(ctx context.Context, bill *models.Bill) error
	GetBills(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Bill, error)
	DeleteBill(ctx context.Context, id string) error

	// Income
	InsertIncome(ctx context.Context, income *models.Income) error
	GetIncome(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Income, error)
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

// Bills
func InsertBill(ctx context.Context, bill *models.Bill) error {
	return implementation.InsertBill(ctx, bill)
}

func GetBills(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Bill, error) {
	return implementation.GetBills(ctx, accountId, date1, date2)
}
func DeleteBill(ctx context.Context, id string) error {
	return implementation.DeleteBill(ctx, id)
}

// Income
func InsertIncome(ctx context.Context, income *models.Income) error {
	return implementation.InsertIncome(ctx, income)
}

func GetIncome(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Income, error) {
	return implementation.GetIncome(ctx, accountId, date1, date2)
}
func DeleteIncome(ctx context.Context, id string) error {
	return implementation.DeleteIncome(ctx, id)
}
