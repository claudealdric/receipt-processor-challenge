package main

import (
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	t.Run("CreatePointsEntry", func(t *testing.T) {
		store := NewInMemoryStore()
		points := 25
		id, err := store.CreatePointsEntry(points)
		HasNoError(t, err)
		DoesNotEqual(t, id, "")
		Equals(t, store.points[id], points)
	})

	t.Run("GetPoints", func(t *testing.T) {
		store := &InMemoryStore{
			points: map[string]int{
				"1": 10,
				"2": 20,
				"3": 30,
			},
		}

		t.Run("returns the points of the given receipt ID", func(t *testing.T) {
			ids := []string{"1", "2", "3"}
			for _, id := range ids {
				got, err := store.GetPoints(id)
				want := store.points[id]
				HasNoError(t, err)
				Equals(t, got, want)
			}
		})

		t.Run("returns a 0 and an error when the given ID does not exist", func(t *testing.T) {
			id := "does-not-exist"
			got, err := store.GetPoints(id)
			want := 0
			HasError(t, err)
			Equals(t, got, want)
		})
	})
}
