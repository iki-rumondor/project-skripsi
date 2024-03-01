package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func AddDate(dateString, amountString string) string {
	amount, _ := strconv.Atoi(amountString)
	if amount < 1 {
		return ""
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	addDate := date.AddDate(0, 0, amount)

	return fmt.Sprintf("%s - %s", date.Format("2006-01-02"), addDate.Format("2006-01-02"))
}

func IsNowInRange(rangeString string) bool {
	parts := strings.Split(rangeString, " - ")
	if len(parts) != 2 {
		return false
	}

	lokasi, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return false
	}

	hariIni := time.Now().In(lokasi)

	start, err := time.Parse("2006-01-02", parts[0])
	if err != nil {
		return false
	}

	end, err := time.Parse("2006-01-02", parts[1])
	if err != nil {
		return false
	}

	return hariIni.After(start) && hariIni.Before(end)

}
