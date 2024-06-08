package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"reciept-processor-fetch/internal"
	"reciept-processor-fetch/models"
)

func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := internal.ProcessReceipt(receipt)
	response := map[string]string{"id": id}
	json.NewEncoder(w).Encode(response)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := internal.GetPoints(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": points}
	json.NewEncoder(w).Encode(response)
}
