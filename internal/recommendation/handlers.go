package recommendation

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Waterfountain10/oppa-suggested-this/pkg/models"
)

type Handlers struct {
	engine *RecEngine
}

func NewHandlers(engine *RecEngine) *Handlers {
	return &Handlers{engine: engine}
}

func (h *Handlers) AddContent(w http.ResponseWriter, r *http.Request) {
	// Only POST allowed
	if r.Method != http.MethodPost {
		http.Error(w, "nope, use POST", http.StatusMethodNotAllowed)
		return
	}

	var content models.Content
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		log.Printf("bad content json: %v", err)
		http.Error(w, "couldn't parse that content", http.StatusBadRequest)
		return
	}

	if err := h.engine.AddContent(content); err != nil {
		log.Printf("failed to add content: %v", err)
		http.Error(w, "something went wrong adding that", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) AddRating(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "nope, use POST", http.StatusMethodNotAllowed)
		return
	}

	var rating models.UserRating
	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		log.Printf("bad rating json: %v", err)
		http.Error(w, "couldn't parse that rating", http.StatusBadRequest)
		return
	}

	if err := h.engine.AddRating(rating); err != nil {
		log.Printf("failed to add rating: %v", err)
		http.Error(w, "something went wrong adding that rating", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) GetRecs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "nope, use GET", http.StatusMethodNotAllowed)
		return
	}

	// Parse query params
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "need a userId", http.StatusBadRequest)
		return
	}

	contentType := r.URL.Query().Get("type")
	if contentType == "" {
		http.Error(w, "need a content type (drama/game/music)", http.StatusBadRequest)
		return
	}

	limit := 10 // Default
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	// Get recommendations
	recs := h.engine.GetRecs(userID, contentType, limit)

	// Build response
	resp := models.RecResponse{
		UserID:          userID,
		Recommendations: recs,
	}

	// Add a fun reason if we found some recs
	if len(recs) > 0 {
		resp.ReasonGiven = "Based on your love for " + recs[0].Genres[0]
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}
