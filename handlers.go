package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
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
