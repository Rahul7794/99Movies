package handler

import (
	"reflect"
	"testing"

	"99Movies/models"
)

func TestInputs_CreateTweets(t *testing.T) {
	tests := []struct {
		name       string
		setup      func(chan models.Reviews)
		checkTweet func(chan string)
		MovieMap   func() map[string]int
		checkError  func(chan error)
		wantError  bool
	}{
		{
			name: "successfully create tweets",
			setup: func(out chan models.Reviews) {
				reviews := models.Reviews{
					Title:  "Star Wars",
					Review: "Great, this film was",
					Score:  77,
				}
				out <- reviews
			},
			MovieMap : func() map[string]int {
				m := make(map[string]int)
				m["Star Wars"] = 1977
				m["Star Wars The Force Awakens"] = 2015
				return m
			},
			checkTweet: func(tweets chan string) {
				for tweet := range tweets {
					expectedTweet := "Star Wars (1977): Great, this film was ****"
					if !reflect.DeepEqual(expectedTweet, tweet) {
						t.Errorf("CreateTweet()=%v, wanted %v", tweet, expectedTweet)
					}
					return
				}
			},
			wantError: false,
		},
		{
			name: "successfully create tweets with year not found",
			setup: func(out chan models.Reviews) {
				reviews := models.Reviews{
					Title:  "Star Wars",
					Review: "Great, this film was",
					Score:  77,
				}
				out <- reviews
			},
			MovieMap : func() map[string]int {
				m := make(map[string]int)
				return m
			},
			checkTweet: func(tweets chan string) {
				for tweet := range tweets {
					expectedTweet := "Star Wars: Great, this film was ****"
					if !reflect.DeepEqual(expectedTweet, tweet) {
						t.Errorf("CreateTweet()=%v, wanted %v", tweet, expectedTweet)
					}
					return
				}
			},
			wantError: false,
		},
		{
			name: "cannot create tweets because of score > 100",
			setup: func(out chan models.Reviews) {
				reviews := models.Reviews{
					Title:  "Star Wars",
					Review: "Great, this film was",
					Score:  101,
				}
				out <- reviews
			},
			MovieMap : func() map[string]int {
				return make(map[string]int)
			},
			checkError: func(errors chan error) {
				for err := range errors {
					expectedErrorString := "cannot get rating from scores the score should be in range 1-100 (inclusive)"
					if !reflect.DeepEqual(err.Error(), expectedErrorString) {
						t.Errorf("CreateTweet()=%v, wanted %v", err, expectedErrorString)
					}
					return
				}
			},
			wantError: true,
		},
	}
	for _, tt := range tests {

		reviewChan := make(chan models.Reviews)
		tweetChan := make(chan string)
		errorChan := make(chan error)
		input := Inputs{
			MoviesMap: tt.MovieMap(),
			TweetChan: tweetChan,
			ErrorChan: errorChan,
			DoneChan:  nil,
			Writer:    nil,
		}
		go tt.setup(reviewChan)
		go input.CreateTweets(reviewChan)
		if tt.wantError {
			tt.checkError(errorChan)
		} else {
			tt.checkTweet(tweetChan)
		}
	}
}

func TestCreateRatings(t *testing.T) {
	tests := []struct {
		name    string
		score   int
		rating  string
		wantErr bool
	}{
		{
			name:    "successfully create ratings",
			score:   65,
			rating:  "***½",
			wantErr: false,
		},
		{
			name:    "cannot get ratings for scores > 100",
			score:   101,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		rating, err := createRatings(tt.score)
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
		if !reflect.DeepEqual(tt.rating, rating) {
			t.Errorf("createRating(%v)=%v, wanted %v", tt.score, rating, tt.rating)
		}
	}
}

func TestTweetsToString(t *testing.T) {
	tests := []struct {
		name   string
		tweets models.Tweet
		output string
	}{
		{
			name: "trim title when the tweets length is more than 140",
			tweets: models.Tweet{
				Title:  "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb",
				Year:   2000,
				Review: "Another movie about a robot. It was a very nice movie and I enjoyed a lot",
				Rating: "****",
			},
			output: "Dr. Strangelove or How I  (2000): Another movie about a robot. It was a very nice movie and I enjoyed a lot ****",
		},
		{
			name: "trim title and review when the tweets length is more than 140",
			tweets: models.Tweet{
				Title:  "Dr. Strangelove or How I Learned to Stop Worrying and Love the Bomb",
				Year:   2000,
				Review: "Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand what was happening because the audio was too low",
				Rating: "****",
			},
			output: "Dr. Strangelove or How I  (2000): Another movie about a robot. Very strong futuristic look. But also very very old. Hard to understand  ****",
		},
		{
			name: "do not trim if the movie size dont increase more than 140",
			tweets: models.Tweet{
				Title:  "Star Wars",
				Year:   2000,
				Review: "Great, this film was",
				Rating: "****",
			},
			output: "Star Wars (2000): Great, this film was ****",
		},
	}

	for _, tt := range tests {
		result := tweetsToString(tt.tweets)
		if !reflect.DeepEqual(tt.output, result) {
			t.Errorf("tweetsToString(%v)=%v, wanted %v", tt.tweets, result, tt.output)
		}
	}
}
