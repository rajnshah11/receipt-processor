package main

import (
	"fmt"
	"net/http"

	"receipt-processor/controllers"
	"receipt-processor/services"
)

func main() {
	service := services.NewReceiptService()
	controller := controllers.NewReceiptController(service)

	http.HandleFunc("/receipts/process", controller.ProcessReceiptHandler)
	http.HandleFunc("/receipts/", controller.PointsHandler)

	fmt.Println("Server running on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
