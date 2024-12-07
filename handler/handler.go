// handler.go
package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"receipt-app/helper"
	"receipt-app/model"
	"strconv"
	"sync"
	"time"
)

var receiptStore = make(map[string]model.Receipt)
var pointsStore = make(map[string]int)
var mu sync.Mutex

// ProcessReceipt handles the /receipts/process POST request.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var receipt model.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	if err := validateReceipt(receipt); err != nil {
		http.Error(w, fmt.Sprintf("Invalid receipt: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Validate each item in the receipt
	for _, item := range receipt.Items {
		if err := validateItem(item); err != nil {
			http.Error(w, fmt.Sprintf("Invalid item: %s", err.Error()), http.StatusBadRequest)
			return
		}
	}

	// Generate a unique ID for the receipt (UUID-like)
	id := fmt.Sprintf("%x", rand.New(rand.NewSource(time.Now().UnixNano())).Int63())

	// Calculate points for the receipt
	points := helper.CalculatePoints(receipt)

	// Store the receipt and the points
	receiptStore[id] = receipt
	pointsStore[id] = points

	// Return the ID in the response
	response := model.ReceiptResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// GetPoints handles the /receipts/{id}/points GET request.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Get the receipt ID from URL
	id := r.URL.Path[len("/receipts/") : len(r.URL.Path)-len("/points")]

	// Fetch the points for the given receipt ID
	points, exists := pointsStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Return the points in the response
	response := model.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ValidateReceipt checks that the required fields are present in the receipt.
func validateReceipt(receipt model.Receipt) error {
	if receipt.Retailer == "" {
		return fmt.Errorf("retailer is required")
	}
	if receipt.PurchaseDate == "" {
		return fmt.Errorf("purchaseDate is required")
	}
	if receipt.PurchaseTime == "" {
		return fmt.Errorf("purchaseTime is required")
	}
	if len(receipt.Items) == 0 {
		return fmt.Errorf("items are required")
	}
	if receipt.Total == "" {
		return fmt.Errorf("total is required")
	}

	// If all required fields are present, return nil
	return nil
}

// ValidateItem checks that each item has a valid short description and price.
func validateItem(item model.Item) error {
	if item.ShortDescription == "" {
		return fmt.Errorf("item short description cannot be empty")
	}

	// Validate price: ensure it's a valid number
	if _, err := strconv.ParseFloat(item.Price, 64); err != nil {
		return fmt.Errorf("item price must be a valid number")
	}

	return nil
}
