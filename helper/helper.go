package helper

import (
	"math"
	"receipt-app/model"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculatePoints calculates the points based on the rules provided.
func CalculatePoints(receipt model.Receipt) int {
	points := 0

	// 1. Points for retailer name
	for _, c := range receipt.Retailer {
		// Check if the character is alphanumeric
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			points++
		}
	}

	// 2. Points for round dollar total
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// 3. Points for multiple of 0.25
	if total*4 == math.Round(total*4) {
		points += 25
	}

	// 4. Points for every 2 items
	points += 5 * (len(receipt.Items) / 2)

	// 5. Points for item description length multiple of 3
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. Points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 7. Points if purchase time is between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
