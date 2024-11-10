package main

import "testing"

func TestCalculateRetailerNamePoints(t *testing.T) {
	tests := []struct {
		retailerName string
		want         int
	}{
		{retailerName: "Target", want: 6},
		{retailerName: "Walmart", want: 7},
		{retailerName: "Giant Eagle", want: 10},
	}

	for _, test := range tests {
		got := calculateRetailerNamePoints(test.retailerName)
		Equals(t, got, test.want)
	}
}
