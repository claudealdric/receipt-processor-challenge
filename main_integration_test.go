package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	store := &InMemoryStore{
		receipts: make(map[string]Receipt),
		points: map[string]int{
			"1": 10,
			"2": 20,
			"3": 30,
		},
	}
	server := NewServer(store)

	t.Run("GET /receipts/{id}/points", func(t *testing.T) {
		t.Run("returns the points of the given, valid receipt ID", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/receipts/1/points", nil)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			HasHttpStatus(t, response.Code, http.StatusOK)
			var pointsResponse PointsResponse
			err := json.NewDecoder(response.Body).Decode(&pointsResponse)
			HasNoError(t, err)
			Equals(t, pointsResponse.Points, 10)
		})

		t.Run("responds with a 404 when given a non-existent ID", func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/receipts/does-not-exist/points", nil)
			response := httptest.NewRecorder()
			server.ServeHTTP(response, request)
			HasHttpStatus(t, response.Code, http.StatusNotFound)
		})

	})
}
