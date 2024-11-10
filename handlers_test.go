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

func TestCalculateItemPoints(t *testing.T) {
	tests := []struct {
		items []ReceiptItem
		want  int
	}{
		{
			items: []ReceiptItem{
				{
					ShortDescription: "Mountain Dew 12PK", // len: 17
					Price:            "6.49",
				},
				{
					ShortDescription: "Emils Cheese Pizza", // len: 18
					Price:            "12.25",
				},
				{
					ShortDescription: "Knorr Creamy Chicken", // len: 20
					Price:            "1.26",
				},
				{
					ShortDescription: "Doritos Nacho Cheese", // len: 20
					Price:            "3.35",
				},
				{
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", // len: 24
					Price:            "12.00",
				},
			},
			want: 16,
		},
		{
			items: []ReceiptItem{
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
			},
			want: 5,
		},
	}

	for _, test := range tests {
		got, err := calculateItemPoints(test.items)
		HasNoError(t, err)
		Equals(t, got, test.want)
	}
}
