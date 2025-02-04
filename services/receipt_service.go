package services

import (
	"math"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"

	"github.com/google/uuid"
)

// ReceiptService handles receipt-related operations.
type ReceiptService struct {
	storage map[string]int
}

// NewReceiptService initializes a new ReceiptService.
func NewReceiptService() *ReceiptService {
	return &ReceiptService{storage: make(map[string]int)}
}

// ProcessReceipt calculates points for a receipt and stores it in memory.
func (s *ReceiptService) ProcessReceipt(receipt models.Receipt) (string, int) {
	points := calculatePoints(receipt)
	id := uuid.New().String()
	s.storage[id] = points
	return id, points
}

// GetPoints retrieves points for a given receipt ID.
func (s *ReceiptService) GetPoints(id string) (int, bool) {
	points, exists := s.storage[id]
	return points, exists
}

// calculatePoints calculates the total points for a given receipt based on the rules.
func calculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
		if isAlphanumeric(char) {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		price, _ := strconv.ParseFloat(item.Price, 64)

		// Rule 5: If item description length is a multiple of 3
		if descLength%3 == 0 {
			points += int(math.Ceil(price * 0.2))
		}
	}

	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)

	// Rule 6: Add 6 points if the day of purchase is odd
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: Add 10 points if purchase time is between 2:00pm and 4:00pm
	if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
		points += 10
	}

	return points
}

// isAlphanumeric checks if a character is alphanumeric (a-z, A-Z, or 0-9).
func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
