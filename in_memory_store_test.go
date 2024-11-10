package main

import (
	"reflect"
	"testing"
)

func TestInMemoryStore(t *testing.T) {
	store := &InMemoryStore{
		receipts: make(map[string]Receipt),
		points: map[string]int{
			"1": 10,
			"2": 20,
			"3": 30,
		},
	}

	t.Run("GetPoints", func(t *testing.T) {
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

func Equals[T any](t testing.TB, got, want T) {
	t.Helper()
	switch v := any(got).(type) {
	case string, int, int64, float64, bool:
		if v != any(want) {
			t.Errorf("got %v, want %v", got, want)
		}
	default:
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func HasError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}
}

func HasNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
