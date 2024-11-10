package main

import (
	"fmt"
	"testing"
)

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

func TestCalculatePurchaseDatePoints(t *testing.T) {
	tests := []struct {
		purchaseDate string
		want         int
	}{
		{purchaseDate: "2024-12-01", want: 6},
		{purchaseDate: "2024-09-10", want: 0},
	}

	for _, test := range tests {
		got, err := calculatePurchaseDatePoints(test.purchaseDate)
		HasNoError(t, err)
		Equals(t, got, test.want)
	}
}

func TestCalculatePurchaseTimePoints(t *testing.T) {
	tests := []struct {
		purchaseTime string
		want         int
	}{
		{purchaseTime: "00:00", want: 0},
		{purchaseTime: "11:59", want: 0},
		{purchaseTime: "12:00", want: 0},
		{purchaseTime: "13:59", want: 0},
		{purchaseTime: "14:00", want: 0},
		{purchaseTime: "14:01", want: 10},
		{purchaseTime: "15:59", want: 10},
		{purchaseTime: "16:00", want: 0},
		{purchaseTime: "23:59", want: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("purchase time: %s", test.purchaseTime), func(t *testing.T) {
			got, err := calculatePurchaseTimePoints(test.purchaseTime)
			HasNoError(t, err)
			Equals(t, got, test.want)
		})
	}
}

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		receipt Receipt
		want    int
	}{
		{
			receipt: Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []ReceiptItem{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					},
					{
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					},
					{
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					},
					{
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					},
					{
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			},
			want: 28,
		},
	}

	for _, test := range tests {
		got, err := CalculatePoints(test.receipt)
		HasNoError(t, err)
		Equals(t, got, test.want)
	}
}
