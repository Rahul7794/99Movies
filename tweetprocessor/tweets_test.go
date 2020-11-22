package tweetprocessor

import (
	"reflect"
	"testing"

	"99Movies/models"
)

func TestInputs_CreateTweets(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Inputs
		tweets    []string
		wantError bool
	}{
		{
			name: "successfully create tweets",
			setup: func() *Inputs {
				m := make(map[string]int)
				m["Star Wars"] = 1977
				m["Star Wars The Force Awakens"] = 2015
				reviews := []models.Reviews{
					{
						Title:  "Star Wars",
						Review: "Great, this film was",
						Score:  77,
					},
					{
						Title:  "Star Wars The Force Awakens",
						Review: "A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some chap came along and basically made the same movie again",
						Score:  50,
					},
				}
				return &Inputs{
					Movies:  m,
					Reviews: reviews,
				}
			},
			tweets: []string{
				"Star Wars (1977): Great, this film was ****",
				"Star Wars The Force Awake (2015): A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some cha **½",
			},
			wantError: false,
		},
		{
			name: "successfully create tweets with year not found",
			setup: func() *Inputs {
				m := make(map[string]int)
				m["Star Wars The Force Awakens"] = 2015
				reviews := []models.Reviews{
					{
						Title:  "Star Wars",
						Review: "Great, this film was",
						Score:  77,
					},
					{
						Title:  "Star Wars The Force Awakens",
						Review: "A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some chap came along and basically made the same movie again",
						Score:  50,
					},
				}
				return &Inputs{
					Movies:  m,
					Reviews: reviews,
				}
			},
			tweets: []string{
				"Star Wars: Great, this film was ****",
				"Star Wars The Force Awake (2015): A long time ago in a galaxy far far away someone made the best sci-fi film of all time. Then some cha **½",
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		input := tt.setup()
		tweets, err := input.CreateTweets()
		if err != nil && tt.wantError == false {
			t.Errorf("unexpected error %s", err)
			return
		}
		if tt.wantError == true && err == nil {
			t.Errorf("wanted an error during %s", tt.name)
			return
		}
		if tt.wantError == true && err != nil {
			// :)
			return
		}
		if !reflect.DeepEqual(tt.tweets, tweets) {
			t.Errorf("CreateTweets=%v, wanted %v", tweets, tt.tweets)
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
