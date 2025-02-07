package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/utils"
)

var receiptStorage = make(map[string]models.Receipt) // In-memory storage

// ProcessReceiptHandler handles the receipt processing endpoint (POST /receipts/process)
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receiptData models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receiptData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	fmt.Println("Generating receiptID for the given receipt")
	receiptID := utils.GenerateID()
	fmt.Println("receiptID generated:", receiptID)
	// Save receipt data
	
	receiptStorage[receiptID] = receiptData

	// Respond with the generated ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}
