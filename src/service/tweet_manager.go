package service

import (
	"github.com/TallerGo/src/domain"
)

var tweet *domain.Tweet

//PublishTweet publish a tweet
func PublishTweet(aTweet *domain.Tweet) {
	tweet = aTweet
}

//GetTweet return the tweet
func GetTweet() *domain.Tweet {
	return tweet
}

//ClearTweet clear a tweet
func ClearTweet() {
	tweet = nil
}
