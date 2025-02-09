package models

// Content represents anything we can recommend - dramas, games, or music
type Content struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Type   string   `json:"type"` // drama|game|music
	Genres []string `json:"genres"`

	// Overall rating (1-10). Using MAL/MDL style rating system
	Rating float64 `json:"rating"`

	// Optional but helpful fields
	ReleaseYear int      `json:"releaseYear,omitempty"`
	Tags        []string `json:"tags,omitempty"`

	// Drama-specific fields
	Episodes   int    `json:"episodes,omitempty"`
	StudioName string `json:"studioName,omitempty"`

	// Game-specific fields
	Platform  string `json:"platform,omitempty"`
	Publisher string `json:"publisher,omitempty"`

	// Music-specific fields
	Artist string `json:"artist,omitempty"`
	Album  string `json:"album,omitempty"`
}

// UserRating captures what a user thought about a piece of content
type UserRating struct {
	UserID    string  `json:"userId"`
	ContentID string  `json:"contentId"`
	Score     float64 `json:"score"`   // 1-10 scale
	Dropped   bool    `json:"dropped"` // Did they drop it?
	Completed bool    `json:"completed"`
	Notes     string  `json:"notes,omitempty"`
}

// RecRequest bundles up what we need to make recommendations
type RecRequest struct {
	UserID      string `json:"userId"`
	ContentType string `json:"type"`  // What kind of content they want
	MaxResults  int    `json:"limit"` // How many recs to return (default 10)
}

// RecResponse wraps up our recommendations with some context
type RecResponse struct {
	UserID          string    `json:"userId"`
	Recommendations []Content `json:"recommendations"`
	ReasonGiven     string    `json:"reason,omitempty"` // Why we recommended these
}
