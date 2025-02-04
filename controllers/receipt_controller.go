package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"receipt-processor/models"
	"receipt-processor/services"
)

// ReceiptController handles HTTP requests related to receipts.
type ReceiptController struct {
	service *services.ReceiptService
}

// NewReceiptController initializes a new ReceiptController.
func NewReceiptController(service *services.ReceiptService) *ReceiptController {
	return &ReceiptController{service: service}
}

// ProcessReceiptHandler handles POST requests to /receipts/process.
func (c *ReceiptController) ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil || !validateReceipt(receipt) {
		http.Error(w, "Invalid JSON payload or missing fields", http.StatusBadRequest)
		return
	}

	id, _ := c.service.ProcessReceipt(receipt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ProcessResponse{ID: id})
}

// PointsHandler handles GET requests to /receipts/{id}/points.
func (c *ReceiptController) PointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimSuffix(strings.TrimPrefix(r.URL.Path, "/receipts/"), "/points")
	points, exists := c.service.GetPoints(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.PointsResponse{Points: points})
}

// validateReceipt ensures all required fields in a receipt are present.
func validateReceipt(receipt models.Receipt) bool {
	return receipt.Retailer != "" && receipt.PurchaseDate != "" &&
		receipt.PurchaseTime != "" && receipt.Total != ""
}
