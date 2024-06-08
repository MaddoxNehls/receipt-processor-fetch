package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"reciept-processor-fetch/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
