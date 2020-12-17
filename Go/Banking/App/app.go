package App

import (
	"banking/Domain"
	"banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(Domain.NewCustomerRepositoryDB())}

	router.HandleFunc("/customers", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
