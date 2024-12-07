// model.go
package model

// Receipt represents the structure of a receipt.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents each individual item in the receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ReceiptResponse represents the response when processing a receipt.
type ReceiptResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response with the number of points.
type PointsResponse struct {
	Points int `json:"points"`
}
