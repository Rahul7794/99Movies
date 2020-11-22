package tweetprocessor

import (
	"fmt"

	"99Movies/models"
)

const (
	maxTweetLength = 140
	maxTitleLength = 25
)

// CreateTweets creates array of tweets from movies and reviews json
func (input *Inputs) CreateTweets() ([]string, error) {
	var tweets []string
	for i := 0; i < len(input.Reviews); i++ {
		// Get rating based on scores
		ratings, err := createRatings(input.Reviews[i].Score)
		if err != nil {
			return nil, fmt.Errorf("cannot get rating from scores %v", err)
		}
		tweet := models.Tweet{
			Title:  input.Reviews[i].Title,
			Year:   input.getYear(input.Reviews[i].Title),
			Review: input.Reviews[i].Review,
			Rating: ratings,
		}
		tweets = append(tweets, tweetsToString(tweet))
	}
	return tweets, nil
}

// getYear get year models.Movies struct and delete entry from map.
func (input *Inputs) getYear(movies string) int {
	if year, found := input.Movies[movies]; found {
		delete(input.Movies, movies)
		return year
	}
	return 0
}

// createRatings convert scores to weighted 5 star rating and half
func createRatings(scores int) (string, error) {
	switch {
	case inBetween(scores, 1, 18):
		return "½", nil
	case inBetween(scores, 19, 27):
		return "*", nil
	case inBetween(scores, 28, 36):
		return "*½", nil
	case inBetween(scores, 37, 45):
		return "**", nil
	case inBetween(scores, 46, 54):
		return "**½", nil
	case inBetween(scores, 55, 63):
		return "***", nil
	case inBetween(scores, 64, 72):
		return "***½", nil
	case inBetween(scores, 73, 81):
		return "****", nil
	case inBetween(scores, 82, 90):
		return "****½", nil
	case inBetween(scores, 91, 100):
		return "*****", nil
	default:
		return "", fmt.Errorf("the score should be in range 1-100 (inclusive)")
	}
}

// inBetween check if the val falls in range(low, high) inclusive
func inBetween(val, low, high int) bool {
	return low <= val && val <= high
}

// tweetsToString convert models.Tweet to string for output
func tweetsToString(tweet models.Tweet) string {
	tweetString := tweet.String()
	// check for the length of tweet, if greater than maxlength then trim it.
	if len(tweetString) > maxTweetLength {
		return trimReview(tweet)
	}
	return tweetString
}

// trimReview trims the tweets if its exceeds maxTweetLength
func trimReview(tweet models.Tweet) string {
	// if length of title is greater than maxTitleLength then trim the title to maxTitleLength
	if len(tweet.Title) > maxTitleLength {
		tweet.Title = tweet.Title[:maxTitleLength]
	}
	// convert to string to check if it exceeds the maxlength even after trimming title
	tweetString := tweet.String()

	if len(tweetString) > maxTweetLength {
		// if its exceeds, trim reviews to keep the tweet string length to maxTweetLength
		remainingLengthToTrim := len(tweetString) - maxTweetLength
		tweet.Review = tweet.Review[:len(tweet.Review)-remainingLengthToTrim]
	}
	return tweet.String()
}
