package main

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type PointsResponse struct {
	Points int `json:"points"`
}

func (s *Server) HandleGetPoints(w http.ResponseWriter, r *http.Request) {
	receiptId := r.PathValue("id")
	points, err := s.store.GetPoints(receiptId)
	if err != nil {
		http.Error(w, "no receipt found for the given ID", http.StatusNotFound)
		return
	}

	response := PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func CalculatePoints(receipt Receipt) (int, error) {
	var points int

	retailerNamePoints := calculateRetailerNamePoints(receipt.Retailer)
	points += retailerNamePoints

	dollarTotalPoints, err := calculateDollarTotalPoints(receipt.Total)
	if err != nil {
		return 0, err
	}
	points += dollarTotalPoints

	itemPoints, err := calculateItemPoints(receipt.Items)
	if err != nil {
		return 0, err
	}
	points += itemPoints

	purchaseDatePoints, err := calculatePurchaseDatePoints(receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	points += purchaseDatePoints

	purchaseTimePoints, err := calculatePurchaseTimePoints(receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	points += purchaseTimePoints

	return points, nil
}

func calculateRetailerNamePoints(name string) int {
	// Rule 1: one point for every alphanumeric character in the retailer name
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	matches := alphanumeric.FindAllString(name, -1)
	return len(matches)
}

func calculateDollarTotalPoints(total string) (int, error) {
	var points int

	amount, err := strconv.ParseFloat(total, 32)
	if err != nil {
		return 0, err
	}

	amountInCents := int(amount * 100)

	// Rule 2: 50 points if the total is a round dollar amount with no cents
	if amountInCents%100 == 0 {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if amountInCents%25 == 0 {
		points += 25
	}

	return points, nil
}

func calculateItemPoints(items []ReceiptItem) (int, error) {
	var points int

	// Rule 4: 5 points for every two items on the receipt
	itemCountPoints := len(items) / 2 * 5
	points += itemCountPoints

	// Rule 5: if the trimmed length of the description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 != 0 {
			continue
		}
		price, err := strconv.ParseFloat(item.Price, 32)
		if err != nil {
			continue
		}
		itemDescriptionPoints := int(math.Ceil(price * 0.2))
		points += itemDescriptionPoints
	}

	return points, nil
}

func calculatePurchaseDatePoints(purchaseDate string) (int, error) {
	// Rule 6: 6 points if the day in the purchase day is odd
	date, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return 0, err
	}

	if date.Day()%2 == 1 {
		return 6, nil
	}

	return 0, nil
}

func calculatePurchaseTimePoints(purchaseTime string) (int, error) {
	// Rule 7: 10 points if the time of purchase is after 2:00pm and before
	// 4:00pm
	timeVal, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0, err
	}
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")

	purchaseDateTime := time.Date(0, 1, 1, timeVal.Hour(), timeVal.Minute(), 0, 0, time.UTC)
	startDateTime := time.Date(0, 1, 1, startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
	endDateTime := time.Date(0, 1, 1, endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)

	if purchaseDateTime.After(startDateTime) && purchaseDateTime.Before(endDateTime) {
		return 10, nil
	}

	return 0, nil
}
