package App

import (
	"banking/Domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(Domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customers", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDR")
	port := os.Getenv("SERVER_PORT")

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
