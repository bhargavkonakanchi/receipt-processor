package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-processor/handlers"
)

func main() {
	// Routes for receipt processing
	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	http.HandleFunc("/receipts/", handlers.GetPointsHandler)

	// Start the server
	fmt.Println("Server started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
