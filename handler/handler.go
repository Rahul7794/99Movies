package handler

import (
	"99Movies/io"
	"99Movies/models"
)

// Processor contains methods to be implemented when create tweets
type Processor interface {
	CreateTweets(chan models.Reviews) // CreateTweets creates tweets
	SaveTweets()                      // SaveTweets save it to output source provided.
}

// Inputs contains movies and reviews to create tweets
type Inputs struct {
	MoviesMap map[string]int     // MoviesMap stores release date of movies
	TweetChan chan string        // TweetChan stores tweet produced
	ErrorChan chan<- error       // ErrorChan stores errors if any
	DoneChan  chan<- bool        // DoneChan indicates completion of process
	Writer    io.WriterInterface // Writer interface contain write method
}

// New initialize object
func New(input Inputs) Processor {
	return &Inputs{
		MoviesMap: input.MoviesMap,
		TweetChan: input.TweetChan,
		ErrorChan: input.ErrorChan,
		DoneChan:  input.DoneChan,
		Writer:    input.Writer,
	}
}
