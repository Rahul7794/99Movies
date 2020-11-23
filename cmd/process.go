package cmd

import (
	"github.com/spf13/cobra"

	"99Movies/handler"
	"99Movies/io"
	"99Movies/log"
	"99Movies/models"
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

	// Initialize all the channels used in application.
	reviewsChan := make(chan models.Reviews)
	errorChan := make(chan error)
	tweetChan := make(chan string)
	doneChan := make(chan bool, 1)

	// Create a Movies reader object.
	movies := io.NewReader("file", moviesPath)
	defer movies.Close()
	// Create a Reviews reader object.
	readerReviews := io.NewReader("file", reviewsPath)
	defer readerReviews.Close()

	// Create a tweet writer object
	writerTweets := io.NewWriter("console", nil)

	// create Map of movies and date
	moviesMap, err := movies.GetMovies()
	if err != nil {
		return err
	}
	// read reviews
	go readerReviews.GetReviews(reviewsChan, errorChan)

	// create input for processor
	input := handler.Inputs{
		MoviesMap: moviesMap,
		TweetChan: tweetChan,
		ErrorChan: errorChan,
		DoneChan:  doneChan,
		Writer:    writerTweets,
	}

	// create a processor object
	tweetsProcessor := handler.New(input)

	// compose tweets
	go tweetsProcessor.CreateTweets(reviewsChan)

	// save tweets
	go tweetsProcessor.SaveTweets()

	// Listen to error and done channel
	// If error channel receive error, return it
	// if done channel receive signal, end the program
	for {
		select {
		case err := <-errorChan: // Listen errorChannel
			return err
		case <-doneChan: // Listen done Channel
			log.Info("complete !!!")
			return nil
		}
	}
}
