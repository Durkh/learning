package App

import (
	service2 "banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service2.CustomerService
}

func (ch *CustomerHandler) getCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(w, err.ID, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	id := args["customer_id"]

	customer, err := ch.service.GetCustomerByID(id)
	if err != nil {
		writeResponse(w, err.ID, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
