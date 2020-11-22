package io

import (
	"encoding/json"
	"os"

	"99Movies/log"
	"99Movies/models"
)

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

// JSONToMovies reads Movie Json and store it in map
func (reader *JSONReader) JSONToMovies() (map[string]int, error) {
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

// JSONToReviews reads reviews Json and store it in arrays
func (reader *JSONReader) JSONToReviews() ([]models.Reviews, error) {
	var result []models.Reviews
	for reader.HasNext() {
		err := reader.ParseFromJSON(&result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// NewFileReader takes filename as input and creates an ReaderInterface object
func NewFileReader(filename string) ReaderInterface {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("cannot open the file %v", err)
	}
	decoder := json.NewDecoder(file)
	return &JSONReader{
		Parser: decoder,
		File:   file,
	}
}
