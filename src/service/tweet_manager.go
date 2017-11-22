package service

import (
	"fmt"

	"github.com/TallerGo/src/domain"
)

var tweet *domain.Tweet

//PublishTweet publish a tweet
func PublishTweet(aTweet *domain.Tweet) error {
	if aTweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if aTweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(aTweet.Text) > 140 {
		return fmt.Errorf("text exceeding 140 characters")
	}
	tweet = aTweet
	return nil
}

//GetTweet return the tweet
func GetTweet() *domain.Tweet {
	return tweet
}

//ClearTweet clear a tweet
func ClearTweet() {
	tweet = nil
}
