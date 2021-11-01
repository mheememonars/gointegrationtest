package main

import (
	"go_integration_test/config"
	"go_integration_test/customer"
	"go_integration_test/db"
	"net/http"
	"os"
)

func init() {

	config.Setup(os.Getenv("ENV"), "./env")

	db.Init()
}

func main() {
	http.HandleFunc("/customer/add", customer.CustomerAddHandler)
	http.HandleFunc("/customer/find", customer.CustomerFindHandler)

	http.ListenAndServe(":8080", nil)
}
