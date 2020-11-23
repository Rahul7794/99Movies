package io

import (
	"99Movies/models"
)

// ReaderInterface provides an abstraction over Reader
type ReaderInterface interface {
	GetMovies() (map[string]int, error)
	GetReviews(chan<- models.Reviews, chan<- error)
	Close() error
}

// WriterInterface provides abstraction over writer
type WriterInterface interface {
	WriteTweets(tweet string)
}

// InputSource store key, value pair, creating object for different input source
var InputSource = map[string]func(in interface{}) ReaderInterface{
	"file": NewFileReader,
}

// OutputSource store key, value pair, creating object for different output source
var OutputSource = map[string]func(out interface{}) WriterInterface{
	"console": NewConsoleWriter,
}

// NewReader returns object of ReaderInterface
func NewReader(source string, in interface{}) ReaderInterface {
	if _, ok := InputSource[source]; ok {
		return InputSource[source](in)
	}
	return nil
}

// NewWriter returns object of WriterInterface
func NewWriter(source string, out interface{}) WriterInterface {
	if _, ok := OutputSource[source]; ok {
		return OutputSource[source](out)
	}
	return nil
}
