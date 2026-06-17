package utils

import (
	"time"
)

func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	birthdayThisYear := time.Date(now.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)
	if now.Before(birthdayThisYear) {
		age--
	}

	return age
}