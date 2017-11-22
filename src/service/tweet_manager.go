package service

import (
	"fmt"

	"github.com/TallerGo/src/domain"
)

var tweet *domain.Tweet

//PublishTweet publish a tweet
func PublishTweet(aTweet *domain.Tweet) error {
	var err error
	if aTweet.User != "" {
		tweet = aTweet
	} else {
		err = fmt.Errorf("user is required")
	}
	return err
}

//GetTweet return the tweet
func GetTweet() *domain.Tweet {
	return tweet
}

//ClearTweet clear a tweet
func ClearTweet() {
	tweet = nil
}
