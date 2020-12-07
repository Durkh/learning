package App

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Clients struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getClients(w http.ResponseWriter, r *http.Request) {
	customers := []Clients{
		{"Egidio", "Patos", "58000"},
		{"Rondon", "Patos", "58000"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}