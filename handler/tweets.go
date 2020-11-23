package handler

import (
	"fmt"

	"99Movies/models"
)

const (
	maxTweetLength = 140
	maxTitleLength = 25
)

// CreateTweets create tweets from movies and reviews by user
func (i *Inputs) CreateTweets(in chan models.Reviews) {
	for review := range in {
		// Get rating based on scores
		ratings, err := createRatings(review.Score)
		if err != nil {
			i.ErrorChan <- fmt.Errorf("cannot get rating from scores %v", err)
		}
		tweet := models.Tweet{
			Title:  review.Title,
			Year:   i.getYear(review.Title),
			Review: review.Review,
			Rating: ratings,
		}
		i.TweetChan <- tweetsToString(tweet)
	}
	close(i.TweetChan)
}

// SaveTweets write tweets to different output source
func (i *Inputs) SaveTweets() {
	for t := range i.TweetChan {
		i.Writer.WriteTweets(t)
	}
	i.DoneChan <- true
}

// getYear get year models.Movies struct and delete entry from map.
func (i *Inputs) getYear(movies string) int {
	if year, found := i.MoviesMap[movies]; found {
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
