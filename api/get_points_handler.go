package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/claudealdric/receipt-processor-challenge/data"
)

type PointsResponse struct {
	Points int `json:"points"`
}

func (s *Server) HandleGetPoints(w http.ResponseWriter, r *http.Request) {
	receiptId := r.PathValue("id")
	points, err := s.store.GetPoints(receiptId)
	if err != nil {
		if errors.Is(err, data.ErrReceiptNotFound) {
			http.Error(w, "no receipt found for the given ID", http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response := PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
