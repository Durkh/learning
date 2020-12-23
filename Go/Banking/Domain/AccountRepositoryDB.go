package Domain

import (
	"banking/errs"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		log.Println("Error creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting ID from new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Database error")
	}

	account.AccountID = strconv.FormatInt(id, 10)

	return &account, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
