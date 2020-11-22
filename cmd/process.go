package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"99Movies/io"
	"99Movies/log"
	"99Movies/tweetprocessor"
	"99Movies/version"
)

// Command line Args
var runCmd = &cobra.Command{
	Use:   "process", // SubCommand
	Short: "compose tweets for movies to share",
	Long:  "read in movie reviews that employees have written, and then compose tweets that can be shared through company account",
	RunE:  processCmd, // compose movies tweets
}

// reviewsPath for file of movies reviews.
var reviewsPath string

// moviesPath for file of movies information
var moviesPath string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&reviewsPath, "reviews", "r", "", "input path for the json file")
	err := runCmd.MarkFlagRequired("reviews")
	handleCommandError(err)
	runCmd.Flags().StringVarP(&moviesPath, "movies", "m", "", "input path for the json file")
	err = runCmd.MarkFlagRequired("movies")
	handleCommandError(err)
}

func processCmd(_ *cobra.Command, _ []string) error {
	log.Infof("starting the tweet processor with version: %s %s \n on %s => %s", version.Version, version.BuildDate,
		version.OsArch, version.GoVersion)

	// Create a Movies reader object.
	readerMovies := io.NewReader("fileType", moviesPath)
	defer readerMovies.Close()
	// Create a Reviews reader object.
	readerReviews := io.NewReader("fileType", reviewsPath)
	defer readerReviews.Close()

	// read movies
	movies, err := readerMovies.JSONToMovies()
	if err != nil {
		return err
	}
	// read reviews
	reviews, err := readerReviews.JSONToReviews()
	if err != nil {
		return err
	}

	// create input for processor
	input := tweetprocessor.Inputs{
		Movies:  movies,
		Reviews: reviews,
	}

	// create a processor object
	tweetsProcessor := tweetprocessor.New(input)

	// compose tweets
	tweets, err := tweetsProcessor.CreateTweets()
	if err != nil {
		return err
	}

	// print tweets to console
	// todo: need to be properly implemented to send tweets to specific output destination.
	for _, tweets := range tweets {
		fmt.Println(tweets)
	}
	return nil
}
