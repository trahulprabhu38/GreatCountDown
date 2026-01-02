package handlers

import (
	"countdown/utils"
	"encoding/json"
	"net/http"
)

// GetYearCountdown handles GET /api/countdown requests
func GetYearCountdown(w http.ResponseWriter, r *http.Request) {
	// Get year 2026 progress
	data := utils.GetYearProgress(2026)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	json.NewEncoder(w).Encode(data)
}

// GetMonthCountdown handles GET /api/countdown/month requests
func GetMonthCountdown(w http.ResponseWriter, r *http.Request) {
	// Get current month progress
	data := utils.GetMonthProgress()

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	json.NewEncoder(w).Encode(data)
}

// GetWeekCountdown handles GET /api/countdown/week requests
func GetWeekCountdown(w http.ResponseWriter, r *http.Request) {
	// Get current week progress
	data := utils.GetWeekProgress()

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send response
	json.NewEncoder(w).Encode(data)
}

// CORSMiddleware adds CORS headers for development
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from Vite dev server
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
