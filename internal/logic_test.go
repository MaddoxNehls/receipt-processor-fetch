package internal

import (
	"testing"
	"reciept-processor-fetch/models"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Multiple Items",
			receipt: models.Receipt{
				Retailer: "Walmart",
				PurchaseDate: "2023-05-15",
				PurchaseTime: "15:45",
				Items: []models.Item{
					{ShortDescription: "Bread - 1 loaf", Price: "1.25"},
					{ShortDescription: "Milk - 1 gallon", Price: "3.99"},
					{ShortDescription: "Eggs - 12 count", Price: "2.50"},
					{ShortDescription: "Cheese - 1 block", Price: "5.00"},
					{ShortDescription: "Chicken - 1 lb", Price: "8.01"},
				},
				Total: "20.75",
			},
			expected: 50,
		},
		{
			name: "Round Dollar Amount",
			receipt: models.Receipt{
				Retailer: "Best Buy",
				PurchaseDate: "2022-08-20",
				PurchaseTime: "14:30",
				Items: []models.Item{
					{ShortDescription: "USB Cable", Price: "10.00"},
					{ShortDescription: "Mouse Pad", Price: "20.00"},
					{ShortDescription: "Keyboard", Price: "20.00"},
				},
				Total: "50.00",
			},
			expected: 103,
		},
		{
			name: "Single Item Specific Length",
			receipt: models.Receipt{
				Retailer: "Apple Store",
				PurchaseDate: "2022-12-13",
				PurchaseTime: "16:00",
				Items: []models.Item{
					{ShortDescription: "iPhone 12 Pro Max", Price: "999.99"},
				},
				Total: "999.99",
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := CalculatePoints(tt.receipt)
			if points != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, points)
			}
		})
	}
}
