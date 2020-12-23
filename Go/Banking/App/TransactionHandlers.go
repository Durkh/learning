package App

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (t TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["account_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountID = id
		transaction, appErr := t.service.NewTransaction(request)
		if appErr != nil {
			writeResponse(w, appErr.ID, appErr.Message)
		} else {
			writeResponse(w, http.StatusOK, transaction)
		}
	}
}
