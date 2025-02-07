package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-processor/utils"
	"strings"
)

// GetPointsHandler handles the /receipts/{id}/points endpoint (GET)
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	receiptID := pathParts[2]

	// Look up the receipt
	receiptData, found := receiptStorage[receiptID]
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points for the receipt
	fmt.Println("Calculating points for receiptID:", receiptID)
	points := utils.CalculatePoints(receiptData)
	fmt.Println("points for receiptID", receiptID, "=", points)
	// Respond with the calculated points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
