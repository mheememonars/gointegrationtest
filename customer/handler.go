package customer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json"
)

func CustomerAddHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CustomerAddHandler")
	if r.Method == "POST" {
		r.ParseForm()
		w.Header().Set(contentTypeKey, contentTypeValue)
		var requestAddCustomer CustomerData
		requestData, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(requestData, &requestAddCustomer)
		if err == nil {
			serviceData, errService := AddCustomer(requestAddCustomer)
			if errService == nil {
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(serviceData)
			}
		}
	}
	return
}

func CustomerFindHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CustomerFindHandler")
	if r.Method == "POST" {
		r.ParseForm()
		w.Header().Set(contentTypeKey, contentTypeValue)
		var requestFindCustomer RequestFindCustomer
		requestData, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(requestData, &requestFindCustomer)
		if err == nil {
			serviceData, errService := FindCustomer(requestFindCustomer)
			if errService == nil {
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(serviceData)
			}
		}
	}
	return
}
