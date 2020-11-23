package io

import (
	"encoding/json"
	"fmt"
	"os"

	"99Movies/log"
	"99Movies/models"
)

const delimit = '['

// JSONReader type has json.Decoder as field
type JSONReader struct {
	Parser *json.Decoder
	File   *os.File
}

// ParseFromJSON decodes json to struct provided
func (reader *JSONReader) ParseFromJSON(c interface{}) error {
	return reader.Parser.Decode(c)
}

// HasNext checks if there is any more elements to be decoded
func (reader *JSONReader) HasNext() bool {
	return reader.Parser.More()
}

// Close closes the File, rendering it unusable for I/O.
func (reader *JSONReader) Close() error {
	return reader.File.Close()
}

// GetMovies reads Movie Json and store it in map
func (reader *JSONReader) GetMovies() (map[string]int, error) {
	result := make(map[string]int)
	var movies []models.Movies
	for reader.HasNext() {
		err := reader.ParseFromJSON(&movies)
		if err != nil {
			return nil, err
		}
	}
	for _, movie := range movies {
		if _, found := result[movie.Title]; !found {
			result[movie.Title] = movie.Year
		}
	}
	return result, nil
}

// GetReviews reads reviews Json and store it in arrays
func (reader *JSONReader) GetReviews(out chan<- models.Reviews, error chan<- error) {
	t, err := reader.Parser.Token()
	if err != nil {
		error <- fmt.Errorf("cannot retreive token %v", err)
	}
	tok, ok := t.(json.Delim)
	if ok {
		if tok == delimit {
			for reader.HasNext() {
				var review models.Reviews
				err = reader.ParseFromJSON(&review)
				if err != nil {
					error <- fmt.Errorf("could not decode reviews %v", err)
				}
				out <- review
			}
			close(out)
		}
	}
}

// NewFileReader takes filename as input and creates an ReaderInterface object
func NewFileReader(in interface{}) ReaderInterface {
	file, err := os.Open(in.(string))
	if err != nil {
		log.Fatalf("cannot open the file %v", err)
	}
	decoder := json.NewDecoder(file)
	return &JSONReader{
		Parser: decoder,
		File:   file,
	}
}
