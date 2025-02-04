package models

// Receipt represents the structure of a receipt as defined in the API specification.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents an individual item in a receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// ProcessResponse represents the response for the /receipts/process endpoint.
type ProcessResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response for the /receipts/{id}/points endpoint.
type PointsResponse struct {
	Points int `json:"points"`
}
