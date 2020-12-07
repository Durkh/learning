package App

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start(){

	router := mux.NewRouter()

	router.HandleFunc("/greeting", greet).Methods(http.MethodGet)

	router.HandleFunc("/customers", getClients).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getClientID).Methods(http.MethodGet)


	log.Fatal(http.ListenAndServe("localhost:8000", router))

}

func getClientID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}
