package internal

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"reciept-processor-fetch/models"
)

var receipts = make(map[string]int)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

	// 50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Points for item description length.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() == 14 {
		points += 10
	}

	return points
}

func ProcessReceipt(receipt models.Receipt) string {
	id := uuid.New().String()
	points := CalculatePoints(receipt)
	receipts[id] = points
	return id
}

func GetPoints(id string) (int, bool) {
	points, exists := receipts[id]
	return points, exists
}
