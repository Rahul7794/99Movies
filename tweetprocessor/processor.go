package tweetprocessor

import (
	"99Movies/models"
)

// Processor contains methods to be implemented when create tweets
type Processor interface {
	CreateTweets() ([]string, error)
}

// Inputs contains movies and reviews to create tweets
type Inputs struct {
	Movies  map[string]int
	Reviews []models.Reviews
}

// New initialize object
func New(input Inputs) *Inputs {
	return &input
}
