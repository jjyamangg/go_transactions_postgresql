package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jjyamangg/go_transactions_postgresql/database"
	"github.com/jjyamangg/go_transactions_postgresql/models"
	"github.com/jjyamangg/go_transactions_postgresql/routes"
)

func main() {

	database.DBConnection()

	database.DB.AutoMigrate(models.Account{}, models.Currency{}, models.TransactionType{}, models.Transaction{})

	route := mux.NewRouter()
	route.HandleFunc("/", routes.HomeHandler).Methods("GET")
	route.HandleFunc("/transaction", routes.CreateTransactionHandler).Methods("POST")
	route.HandleFunc("/transactions", routes.GetTransactionsByDateRange).Methods("GET")

	http.ListenAndServe(":3000", route)

}
