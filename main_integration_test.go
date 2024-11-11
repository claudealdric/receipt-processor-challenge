package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/claudealdric/receipt-processor-challenge/api"
	"github.com/claudealdric/receipt-processor-challenge/assert"
	"github.com/claudealdric/receipt-processor-challenge/data"
	"github.com/claudealdric/receipt-processor-challenge/types"
)

func TestServer(t *testing.T) {
	t.Run("process receipt and get the points", func(t *testing.T) {
		// Create a new store and server
		store := data.NewInMemoryStore()
		server := api.NewServer(store)

		// Create receipt for body
		receipt := types.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []types.ReceiptItem{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
				{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
				{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
				{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
			},
			Total: "35.35",
		}

		// Create request body
		body, err := json.Marshal(receipt)
		assert.HasNoError(t, err)

		// Create request
		processReceiptRequest := httptest.NewRequest(
			http.MethodPost,
			"/receipts/process",
			bytes.NewBuffer(body),
		)
		processReceiptRequest.Header.Set("Content-Type", "application/json")

		// Create the response recorder
		processReceiptRr := httptest.NewRecorder()

		// Process the request
		server.ServeHTTP(processReceiptRr, processReceiptRequest)

		// Check status code
		assert.HasHttpStatus(t, processReceiptRr.Code, http.StatusCreated)

		var processReceiptResponse struct {
			Id string `json:"id"`
		}

		err = json.NewDecoder(processReceiptRr.Body).Decode(&processReceiptResponse)
		assert.HasNoError(t, err)

		if processReceiptResponse.Id == "" {
			t.Fatal("expected a non-empty receipt ID")
		}

		// Make a request to get points
		pointsRequest := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/receipts/%s/points", processReceiptResponse.Id),
			nil,
		)
		pointsRr := httptest.NewRecorder()

		// Process request
		server.ServeHTTP(pointsRr, pointsRequest)

		// Check status code
		assert.HasHttpStatus(t, pointsRr.Code, http.StatusOK)

		var pointsResponse api.PointsResponse

		err = json.NewDecoder(pointsRr.Body).Decode(&pointsResponse)
		assert.HasNoError(t, err)
		assert.Equals(t, pointsResponse.Points, 28)
	})
}
