package Domain

import (
	"banking/errs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var (
		err error
	)

	customers := make([]Customer, 0)
	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		log.Println("Error querying the customer table: " + err.Error())
		return nil, errs.NewNUnexpectedError("Unexpected Database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {

	findClientSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, findClientSQL, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error querying the customer table: " + err.Error())
			return nil, errs.NewNUnexpectedError("Unexpected Database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {

	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	Client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{Client}
}
