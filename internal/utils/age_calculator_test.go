package utils

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday already passed this year",
			dob:      time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
			expected: calculateExpectedAge(1990, 5, 15),
		},
		{
			name:     "Birthday today",
			dob:      time.Now().AddDate(-25, 0, 0),
			expected: 25,
		},
		{
			name:     "Birthday in future this year",
			dob:      time.Date(2000, 12, 25, 0, 0, 0, 0, time.UTC),
			expected: calculateExpectedAge(2000, 12, 25),
		},
		{
			name:     "Leap year birthday",
			dob:      time.Date(2000, 2, 29, 0, 0, 0, 0, time.UTC),
			expected: calculateExpectedAge(2000, 2, 29),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			if age != tt.expected {
				t.Errorf("CalculateAge() = %d, want %d", age, tt.expected)
			}
		})
	}
}

func calculateExpectedAge(year, month, day int) int {
	now := time.Now()
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	age := now.Year() - dob.Year()

	birthdayThisYear := time.Date(now.Year(), dob.Month(), dob.Day(), 0, 0, 0, 0, time.UTC)
	if now.Before(birthdayThisYear) {
		age--
	}
	return age
}

func TestCalculateAge_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		dob  time.Time
	}{
		{
			name: "Future date",
			dob:  time.Now().AddDate(1, 0, 0),
		},
		{
			name: "Zero time",
			dob:  time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			if age < 0 {
				t.Logf("Age is negative: %d", age)
			}
		})
	}
}