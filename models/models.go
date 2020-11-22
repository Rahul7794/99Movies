package models

import (
	"fmt"
)

// Movies represent model of movie
type Movies struct {
	Title string `json:"title"`          // Title of a movie
	Year  int    `json:"year,omitempty"` // Year of movie release
}

// Reviews represent model for a reviews
type Reviews struct {
	Title  string `json:"title"`  // Title of a movie
	Review string `json:"review"` // Review of a movie by customer
	Score  int    `json:"score"`  // Score of a movie out of 100
}

// Tweet represent model of tweet
type Tweet struct {
	Title  string // Title of a movie
	Year   int    // Year of movie released
	Review string // Review of a movie by customer
	Rating string // Rating of a movie out of 5 in unicode
}

// String returns output string for a tweet/**/
func (t *Tweet) String() string {
	var result string
	if t.Year == 0 {
		result = fmt.Sprintf("%s: %s %s", t.Title, t.Review, t.Rating)
	} else {
		result = fmt.Sprintf("%s (%d): %s %s", t.Title, t.Year, t.Review, t.Rating)
	}
	return result
}
