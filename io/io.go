package io

import (
	"99Movies/models"
)

// ReaderInterface provides an abstraction over Reader
type ReaderInterface interface {
	JSONToMovies() (map[string]int, error)
	JSONToReviews() ([]models.Reviews, error)
	Close() error
}

// InputSource store key, value pair, creating object for different input source
var InputSource = map[string]func(input string) ReaderInterface{
	"fileType": NewFileReader,
}

// NewReader returns object of ReaderInterface
func NewReader(source string, input string) ReaderInterface {
	if _, ok := InputSource[source]; ok {
		return InputSource[source](input)
	}
	return nil
}
