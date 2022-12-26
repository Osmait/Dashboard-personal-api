package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/osmait/admin-finanzas/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {

	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (id,name,last_name,email,password,token) VALUES ($1,$2,$3,$4,$5,$6)", user.Id, user.Name, user.LastName, user.Email, user.Password, user.Token)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,name,last_name,email FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,token,confirmed FROM users WHERE token =$1", token)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Token, &user.Confirmed); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) UpdateConfirmed(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE users SET confirmed=$1,token =$2 WHERE id =$3", user.Confirmed, user.Token, user.Id)
	return err
}

func (repo *PostgresRepository) GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,name,last_name,email,password,token,confirmed FROM users WHERE email =$1", email)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.Token, &user.Confirmed); err == nil {

			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

// Account

func (repo *PostgresRepository) InsertAccount(ctx context.Context, account *models.Account) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO account (id,name_account,bank,balance,user_id) VALUES ($1,$2,$3,$4,$5)", account.Id, account.Name, account.Bank, account.Balance, account.User_id)
	return err
}

func (repo *PostgresRepository) GetAcoounts(ctx context.Context, userId string) ([]*models.Account, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id,name_account,bank,balance,user_id FROM account WHERE user_id = $1", userId)

	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()
	var accounts []*models.Account
	for rows.Next() {
		var account = models.Account{}
		if err = rows.Scan(&account.Id, &account.Name, &account.Bank, &account.Balance, &account.User_id); err == nil {
			accounts = append(accounts, &account)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (repo *PostgresRepository) DeleteAccount(ctx context.Context, id string, userId string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM account WHERE id = $1 and user_id = $2", id, userId)
	return err
}

// Bills

func (repo *PostgresRepository) InsertBill(ctx context.Context, bill *models.Bill) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO bill (id,bill_name,bill_description,amount,account_id) VALUES ($1,$2,$3,$4,$5)", bill.Id, bill.BillName, bill.BillDescription, bill.Amount, bill.Account_id)
	return err
}

func (repo *PostgresRepository) GetBills(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Bill, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT id,bill_name,bill_description,amount,account_id,created_at FROM bill WHERE account_id = $1 and created_at BETWEEN $2 and $3 ", accountId, date1, date2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()
	var bills []*models.Bill
	for rows.Next() {
		var bill = models.Bill{}
		if err = rows.Scan(&bill.Id, &bill.BillName, &bill.BillDescription, &bill.Amount, &bill.Account_id, &bill.Created_at); err == nil {
			bills = append(bills, &bill)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bills, nil
}

func (repo *PostgresRepository) DeleteBill(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM bill WHERE id = $1 ", id)
	return err
}

// Income

func (repo *PostgresRepository) InsertIncome(ctx context.Context, income *models.Income) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO income (id,icome_name,icome_description,amount,account_id) VALUES ($1,$2,$3,$4,$5)", income.Id, income.IncomeName, income.IncomeDescription, income.Amount, income.Account_id)
	return err
}

func (repo *PostgresRepository) GetIncome(ctx context.Context, accountId string, date1 string, date2 string) ([]*models.Income, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT id,icome_name,icome_description,amount,account_id,created_at FROM bill WHERE account_id = $1 and created_at BETWEEN $2 and $3 ", accountId, date1, date2)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()
	var incomes []*models.Income
	for rows.Next() {
		var income = models.Income{}
		if err = rows.Scan(&income.Id, &income.IncomeName, &income.IncomeDescription, &income.Amount, &income.Account_id, &income.Created_at); err == nil {
			incomes = append(incomes, &income)
		}

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return incomes, nil
}

func (repo *PostgresRepository) DeleteIncome(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM icome WHERE id = $1 ", id)
	return err
}
