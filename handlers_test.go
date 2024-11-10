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

func TestCalculateDollarTotalPoints(t *testing.T) {
	tests := []struct {
		dollarTotal string
		want        int
	}{
		{dollarTotal: "20.00", want: 75},
		{dollarTotal: "212.75", want: 25},
		{dollarTotal: "1234.56", want: 0},
	}

	for _, test := range tests {
		got, err := calculateDollarTotalPoints(test.dollarTotal)
		HasNoError(t, err)
		Equals(t, got, test.want)
	}
}
