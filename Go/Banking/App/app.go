package App

import (
	"banking/Domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	dbClient := ConnectDB()

	customerRepo := Domain.NewCustomerRepositoryDB(dbClient)
	ch := CustomerHandler{service.NewCustomerService(customerRepo)}

	router.HandleFunc("/customers", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)

	accountRepo := Domain.NewAccountRepositoryDB(dbClient)
	ah := AccountHandler{service.NewAccountService(accountRepo)}

	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	transactionRepo := Domain.NewTransactionRepositoryDB(dbClient)
	th := TransactionHandler{service.NewTransactionService(transactionRepo)}
	router.HandleFunc("/customers/{account_id:[0-9]+}/transaction", th.Transaction).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDR")
	port := os.Getenv("SERVER_PORT")

	log.Println("starting server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func sanityCheck() {
	if os.Getenv("SERVER_ADDR") == "" ||
		os.Getenv("SERVER_PORT") == "" {

		log.Fatal("Server environment variables not set")
	}

	if os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {

		log.Fatal("Database environment variables not set")
	}
}

func ConnectDB() *sqlx.DB {
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

	return Client
}
