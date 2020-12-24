package Domain

import (
	"banking/errs"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
	"strings"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (t TransactionRepositoryDB) Record(transaction Transaction) (*Transaction, *errs.AppError) {

	var (
		result sql.Result
		err    error
	)

	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	if strings.ToLower(transaction.TransactionType) == "withdrawal" {
		result, err = t.client.Exec(sqlInsert, transaction.AccountID, -transaction.Amount, transaction.TransactionType,
			transaction.TransactionDate)
	} else {
		result, err = t.client.Exec(sqlInsert, transaction.AccountID, transaction.Amount, transaction.TransactionType,
			transaction.TransactionDate)
	}

	if err != nil {
		log.Println("Error saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting transaction ID: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	transaction.TransactionID = strconv.FormatInt(id, 10)

	return &transaction, nil
}

func (t TransactionRepositoryDB) GetBalance(accountID string) (float64, *errs.AppError) {

	var account Account

	sqlGet := "SELECT amount FROM accounts WHERE account_id = ?"
	err := t.client.Get(&account, sqlGet, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error getting balance: " + err.Error())
			return 0, errs.NewUnexpectedError("Unexpected Database error")
		}
	}

	return account.Amount, nil
}

func (t TransactionRepositoryDB) UpdateBalance(transaction Transaction) (float64, *errs.AppError) {

	sqlUpdate := "UPDATE accounts SET amount = ? WHERE account_id = ?"

	balance, appErr := t.GetBalance(transaction.AccountID)
	if appErr != nil {
		log.Println("Error updating the balance: " + appErr.Message)
		return 0, errs.NewUnexpectedError("Unexpected Database error")
	}
	newBalance := balance + transaction.Amount

	_, err := t.client.Exec(sqlUpdate, newBalance, transaction.AccountID)
	if err != nil {
		log.Println("Error updating the balance: " + err.Error())
		return 0, errs.NewUnexpectedError("Unexpected Database error")
	}

	return newBalance, nil
}

func (t TransactionRepositoryDB) Rollback(transaction Transaction) *errs.AppError {
	sqlDelete := "DELETE FROM transactions WHERE transaction_id = ?"
	_, err := t.client.Exec(sqlDelete, transaction.TransactionID)
	if err != nil {
		log.Println("Error on rollback: " + err.Error())
		return errs.NewUnexpectedError("Unexpected Database error")
	}

	return nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient}
}
