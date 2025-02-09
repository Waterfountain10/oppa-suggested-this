package recommendation

import (
	"fmt"
	"sort"
	"sync"

	"github.com/Waterfountain10/oppa-suggested-this/pkg/models"
)

// RecEngine is the brain of our recommendation system
type RecEngine struct {
	// Using maps for O(1) lookups - we'll need to be fast
	contentByID   map[string]models.Content
	ratingsByUser map[string][]models.UserRating

	// Protect our data from concurrent access
	mu sync.RWMutex

	// Cache stuff we calculate often
	genreScores map[string]map[string]float64 // user -> genre -> score
	cacheMu     sync.RWMutex
}

func NewRecEngine() *RecEngine {
	return &RecEngine{
		contentByID:   make(map[string]models.Content),
		ratingsByUser: make(map[string][]models.UserRating),
		genreScores:   make(map[string]map[string]float64),
	}
}

// AddContent adds new stuff we can recommend
func (e *RecEngine) AddContent(c models.Content) error {
	if c.ID == "" || c.Title == "" {
		return fmt.Errorf("content needs at least an ID and title")
	}

	e.mu.Lock()
	e.contentByID[c.ID] = c
	e.mu.Unlock()

	return nil
}

// AddRating stores what a user thought about something
func (e *RecEngine) AddRating(r models.UserRating) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Make sure the content exists
	if _, exists := e.contentByID[r.ContentID]; !exists {
		return fmt.Errorf("can't rate nonexistent content: %s", r.ContentID)
	}

	// Update/add the rating
	ratings := e.ratingsByUser[r.UserID]
	for i, existing := range ratings {
		if existing.ContentID == r.ContentID {
			ratings[i] = r
			e.ratingsByUser[r.UserID] = ratings

			// Clear their genre score cache since preferences changed
			e.cacheMu.Lock()
			delete(e.genreScores, r.UserID)
			e.cacheMu.Unlock()

			return nil
		}
	}

	// New rating
	e.ratingsByUser[r.UserID] = append(ratings, r)

	// Clear cache
	e.cacheMu.Lock()
	delete(e.genreScores, r.UserID)
	e.cacheMu.Unlock()

	return nil
}

// GetRecs finds content similar to what the user has liked
func (e *RecEngine) GetRecs(userID string, contentType string, limit int) []models.Content {
	if limit <= 0 {
		limit = 10 // Sane default
	}

	e.mu.RLock()
	defer e.mu.RUnlock()

	// Get their genre preferences (using cache if available)
	genrePrefs := e.getUserGenreScores(userID)

	// Score all content of requested type
	type scoredContent struct {
		content models.Content
		score   float64
	}

	var candidates []scoredContent

	for _, content := range e.contentByID {
		// Skip wrong type or already rated content
		if content.Type != contentType || e.hasRated(userID, content.ID) {
			continue
		}

		// Calculate how well it matches their preferences
		score := 0.0
		for _, genre := range content.Genres {
			score += genrePrefs[genre] * content.Rating // Weight by quality too
		}

		if score > 0 {
			candidates = append(candidates, scoredContent{content, score})
		}
	}

	// Sort by score
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	// Return top N
	result := make([]models.Content, 0, limit)
	for i := 0; i < limit && i < len(candidates); i++ {
		result = append(result, candidates[i].content)
	}

	return result
}

// getUserGenreScores calculates how much a user likes each genre
func (e *RecEngine) getUserGenreScores(userID string) map[string]float64 {
	e.cacheMu.RLock()
	if scores, ok := e.genreScores[userID]; ok {
		e.cacheMu.RUnlock()
		return scores
	}
	e.cacheMu.RUnlock()

	scores := make(map[string]float64)

	// Look at everything they rated highly (7+)
	for _, rating := range e.ratingsByUser[userID] {
		if rating.Score >= 7.0 {
			content := e.contentByID[rating.ContentID]
			for _, genre := range content.Genres {
				scores[genre] += rating.Score - 6.0 // More weight for higher scores
			}
		}
	}

	// Cache it
	e.cacheMu.Lock()
	e.genreScores[userID] = scores
	e.cacheMu.Unlock()

	return scores
}

// hasRated checks if a user has already rated something
func (e *RecEngine) hasRated(userID, contentID string) bool {
	for _, rating := range e.ratingsByUser[userID] {
		if rating.ContentID == contentID {
			return true
		}
	}
	return false
}
