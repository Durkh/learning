package Domain

import (
	"banking/errs"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {

	var (
		rows *sql.Rows
		err  error
	)

	if status == "" {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		log.Println("Error querying the customer table" + err.Error())
		return nil, errs.NewNUnexpectedError("Unexpected Database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.status)
		if err != nil {
			log.Println("Error scanning customers" + err.Error())
			return nil, errs.NewNUnexpectedError("Unexpected Database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {

	findClientSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(findClientSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error querying the customer table" + err.Error())
			return nil, errs.NewNUnexpectedError("Unexpected Database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {

	Client, err := sql.Open("mysql", "root:linuxeamor@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{Client}
}
