package io

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"99Movies/models"
)

func TestJSONReader_JSONToMovies(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *json.Decoder
		output  func() map[string]int
		wantErr bool
	}{
		{
			name: "successfully parse",
			setup: func() *json.Decoder {
				input := []byte(`[
									{"title":"Star Wars","year":1977},
									{"title":"Star Wars The Force Awakens","year":2015}
								 ]`)
				decoder := json.NewDecoder(bytes.NewReader(input))
				return decoder
			},
			output: func() map[string]int {
				m := make(map[string]int)
				m["Star Wars"] = 1977
				m["Star Wars The Force Awakens"] = 2015
				return m
			},
			wantErr: false,
		},
		{
			name: "error while parsing",
			setup: func() *json.Decoder {
				input := []byte(`[
									{"abc"},
									{"def"}
								 ]`)
				decoder := json.NewDecoder(bytes.NewReader(input))
				return decoder
			},
			output:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		decoder := tt.setup()
		reader := JSONReader{
			Parser: decoder,
			File:   nil,
		}
		items, err := reader.JSONToMovies()
		if err != nil && tt.wantErr == false {
			t.Errorf("unexpected error %s", err)
			return
		}
		if tt.wantErr == true && err == nil {
			t.Errorf("wanted an error during %s", tt.name)
			return
		}
		if tt.wantErr == true && err != nil {
			// :)
			return
		}
		if !reflect.DeepEqual(tt.output(), items) {
			t.Errorf("JSONToMovies(%v)=%v, wanted %v", tt.output(), items, tt.output())
		}
	}
}

func TestJSONReader_JSONToReviews(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *json.Decoder
		output  []models.Reviews
		wantErr bool
	}{
		{
			name: "successfully parse",
			setup: func() *json.Decoder {
				input := []byte(`[
									{"title":"Star Wars","review":"Great, this film was","score":77},
                                    {"title":"Star Wars The Force Awakens","review":"A long time ago","score":50}
								 ]`)
				decoder := json.NewDecoder(bytes.NewReader(input))
				return decoder
			},
			output: []models.Reviews{
				{
					Title:  "Star Wars",
					Review: "Great, this film was",
					Score:  77,
				},
				{
					Title:  "Star Wars The Force Awakens",
					Review: "A long time ago",
					Score:  50,
				},
			},
			wantErr: false,
		},
		{
			name: "error while parsing",
			setup: func() *json.Decoder {
				input := []byte(`[
									{"abc"},
									{"def"}
								 ]`)
				decoder := json.NewDecoder(bytes.NewReader(input))
				return decoder
			},
			output:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		decoder := tt.setup()
		reader := JSONReader{
			Parser: decoder,
			File:   nil,
		}
		items, err := reader.JSONToReviews()
		if err != nil && tt.wantErr == false {
			t.Errorf("unexpected error %s", err)
			return
		}
		if tt.wantErr == true && err == nil {
			t.Errorf("wanted an error during %s", tt.name)
			return
		}
		if tt.wantErr == true && err != nil {
			// :)
			return
		}
		if !reflect.DeepEqual(tt.output, items) {
			t.Errorf("JSONToMovies(%v)=%v, wanted %v", tt.output, items, tt.output)
		}
	}
}
