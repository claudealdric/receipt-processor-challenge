package main

import (
	"reflect"
	"testing"
)

func DoesNotEqual[T any](t testing.TB, got, want T) {
	t.Helper()
	switch v := any(got).(type) {
	case string, int, int64, float64, bool:
		if v == any(want) {
			t.Errorf("got %v, want NOT %v", got, want)
		}
	default:
		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want NOT %v", got, want)
		}
	}
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

func HasHttpStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
