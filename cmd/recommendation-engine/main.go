package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Waterfountain10/oppa-suggested-this/internal/recommendation"
)

func main() {
	log.Println("Starting up K-Rec Hub recommendation engine...")

	// Create our recommendation engine
	engine := recommendation.NewRecEngine()
	handlers := recommendation.NewHandlers(engine)

	// Set up routes - keeping it simple for now
	mux := http.NewServeMux()
	mux.HandleFunc("/content", handlers.AddContent)
	mux.HandleFunc("/rating", handlers.AddRating)
	mux.HandleFunc("/recommendations", handlers.GetRecs)

	// Add some basic middleware
	handler := addLogging(addCORS(mux))

	// Figure out what port to use
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default
	}

	// Set up the server with some sane timeouts
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Ready to serve recommendations on port %s!", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server died: %v", err)
	}
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s took %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func addCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
