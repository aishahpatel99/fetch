// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-app/handler"
)

func main() {
	// Register the handlers
	http.HandleFunc("/receipts/process", handler.ProcessReceipt)
	http.HandleFunc("/receipts/", handler.GetPoints)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
