package main

import (
	"countdown/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Create a new HTTP multiplexer
	mux := http.NewServeMux()

	// Register API routes
	mux.HandleFunc("/api/countdown", handlers.GetYearCountdown)
	mux.HandleFunc("/api/countdown/month", handlers.GetMonthCountdown)
	mux.HandleFunc("/api/countdown/week", handlers.GetWeekCountdown)

	// Serve static files from frontend/dist
	// Check if dist directory exists
	distPath := "../frontend/dist"
	if _, err := os.Stat(distPath); os.IsNotExist(err) {
		log.Printf("Warning: Frontend dist directory not found at %s\n", distPath)
		log.Println("API endpoints are available, but frontend files will not be served.")
		log.Println("Run 'npm run build' in the frontend directory to build the frontend.")
	} else {
		// Serve static files
		fs := http.FileServer(http.Dir(distPath))
		mux.Handle("/", fs)
		log.Printf("Serving static files from %s\n", distPath)
	}

	// Wrap with CORS middleware
	handler := handlers.CORSMiddleware(mux)

	// Start server
	port := "8080"
	addr := fmt.Sprintf(":%s", port)

	log.Printf("🚀 Server starting on http://localhost%s\n", addr)
	log.Printf("📡 API endpoints:\n")
	log.Printf("   - GET http://localhost%s/api/countdown\n", addr)
	log.Printf("   - GET http://localhost%s/api/countdown/month\n", addr)
	log.Printf("   - GET http://localhost%s/api/countdown/week\n", addr)

	// Get absolute path for better error messages
	absDistPath, _ := filepath.Abs(distPath)
	log.Printf("📁 Static files: %s\n", absDistPath)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
