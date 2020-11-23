package io

import (
	"fmt"
)

type writer struct {
}

// WriteTweets writes tweets to console
func (w *writer) WriteTweets(tweet string) {
	fmt.Println(tweet)
}

// NewConsoleWriter create a console writer object
func NewConsoleWriter(interface{}) WriterInterface {
	return &writer{}
}
