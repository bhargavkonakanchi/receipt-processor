package utils

import (
	"math"
	"receipt-processor/models"
	"regexp"
	"strconv"
	"strings"
	"time"
	"github.com/google/uuid"
)

// Generate a unique ID using the current timestamp and a random number
func GenerateID() string {
	return uuid.New().String()
}

// CalculatePoints calculates the points for a receipt based on the given rules
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1 point for every alphanumeric character in the retailer name
	points += len(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(receipt.Retailer, ""))

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += 5 * (len(receipt.Items) / 2)

	// Calculate item description-based points
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
		}
	}

	// 6 points if the day in the purchase date is odd
	parsedDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the purchase time is between 2:00pm and 4:00pm
	parsedTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if parsedTime.Hour() >= 14 && parsedTime.Hour() < 16 {
		points += 10
	}

	return points
}
