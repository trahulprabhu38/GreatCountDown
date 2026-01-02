package utils

import (
	"math"
	"time"
)

type CountdownData struct {
	Year            int       `json:"year,omitempty"`
	Month           string    `json:"month,omitempty"`
	MonthNumber     int       `json:"monthNumber,omitempty"`
	WeekNumber      int       `json:"weekNumber,omitempty"`
	DaysTotal       int       `json:"daysTotal"`
	DaysPassed      int       `json:"daysPassed"`
	DaysRemaining   int       `json:"daysRemaining"`
	PercentComplete float64   `json:"percentComplete"`
	CurrentDate     time.Time `json:"currentDate"`
	StartDate       time.Time `json:"startDate,omitempty"`
	EndDate         time.Time `json:"endDate"`
}

// GetYearProgress calculates the countdown progress for a given year
func GetYearProgress(year int) CountdownData {
	now := time.Now()

	// Define start and end of the year
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endOfYear := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	// Calculate days
	daysPassed := int(now.Sub(startOfYear).Hours() / 24)
	totalDays := int(endOfYear.Sub(startOfYear).Hours()/24) + 1
	daysRemaining := totalDays - daysPassed

	// Ensure daysPassed is at least 0 and doesn't exceed totalDays
	if daysPassed < 0 {
		daysPassed = 0
	}
	if daysPassed > totalDays {
		daysPassed = totalDays
	}
	if daysRemaining < 0 {
		daysRemaining = 0
	}

	// Calculate percentage
	percentComplete := calculatePercentage(daysPassed, totalDays)

	return CountdownData{
		Year:            year,
		DaysTotal:       totalDays,
		DaysPassed:      daysPassed,
		DaysRemaining:   daysRemaining,
		PercentComplete: percentComplete,
		CurrentDate:     now,
		EndDate:         endOfYear,
	}
}

// GetMonthProgress calculates the countdown progress for the current month
func GetMonthProgress() CountdownData {
	now := time.Now()

	// Get current month details
	year, month, _ := now.Date()

	// Start and end of current month
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)

	// Calculate days
	daysPassed := int(now.Sub(startOfMonth).Hours() / 24)
	totalDays := int(endOfMonth.Sub(startOfMonth).Hours()/24) + 1
	daysRemaining := totalDays - daysPassed

	// Ensure bounds
	if daysPassed < 0 {
		daysPassed = 0
	}
	if daysPassed > totalDays {
		daysPassed = totalDays
	}
	if daysRemaining < 0 {
		daysRemaining = 0
	}

	// Calculate percentage
	percentComplete := calculatePercentage(daysPassed, totalDays)

	return CountdownData{
		Month:           month.String(),
		MonthNumber:     int(month),
		Year:            year,
		DaysTotal:       totalDays,
		DaysPassed:      daysPassed,
		DaysRemaining:   daysRemaining,
		PercentComplete: percentComplete,
		CurrentDate:     now,
		EndDate:         endOfMonth,
	}
}

// GetWeekProgress calculates the countdown progress for the current week (Monday-Sunday)
func GetWeekProgress() CountdownData {
	now := time.Now()

	// Get the current day of week (0 = Sunday, 1 = Monday, etc.)
	currentWeekday := int(now.Weekday())

	// Convert to ISO 8601 format where Monday = 0
	// In Go: Sunday = 0, Monday = 1, ..., Saturday = 6
	// We want: Monday = 0, Tuesday = 1, ..., Sunday = 6
	isoWeekday := currentWeekday - 1
	if isoWeekday < 0 {
		isoWeekday = 6 // Sunday
	}

	// Calculate start of week (Monday)
	startOfWeek := now.AddDate(0, 0, -isoWeekday)
	startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, time.UTC)

	// Calculate end of week (Sunday)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	endOfWeek = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 23, 59, 59, 0, time.UTC)

	// Get ISO week number
	_, weekNumber := now.ISOWeek()

	// Calculate days (0-indexed: Monday = 0, Sunday = 6)
	daysPassed := isoWeekday + 1 // +1 to include current day
	totalDays := 7
	daysRemaining := totalDays - daysPassed

	// Ensure bounds
	if daysPassed < 0 {
		daysPassed = 0
	}
	if daysPassed > totalDays {
		daysPassed = totalDays
	}
	if daysRemaining < 0 {
		daysRemaining = 0
	}

	// Calculate percentage
	percentComplete := calculatePercentage(daysPassed, totalDays)

	year, _ := now.ISOWeek()

	return CountdownData{
		WeekNumber:      weekNumber,
		Year:            year,
		DaysTotal:       totalDays,
		DaysPassed:      daysPassed,
		DaysRemaining:   daysRemaining,
		PercentComplete: percentComplete,
		CurrentDate:     now,
		StartDate:       startOfWeek,
		EndDate:         endOfWeek,
	}
}

// calculatePercentage calculates the percentage with 2 decimal precision
func calculatePercentage(passed, total int) float64 {
	if total == 0 {
		return 0
	}

	percentage := (float64(passed) / float64(total)) * 100
	// Round to 2 decimal places
	return math.Round(percentage*100) / 100
}
