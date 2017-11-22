package service

import (
	"fmt"

	"github.com/TallerGo/src/domain"
)

var tweets []*domain.Tweet

//PublishTweet publish a tweet
func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	}
	tweets = append(tweets, tweet)
	return nil
}

//GetTweet return the tweet
func GetTweet() *domain.Tweet {
	if len(tweets) != 0 {
		return tweets[len(tweets)-1]
	}
	return nil
}

//ClearTweet clear a tweet
func ClearTweet() {
	tweets = tweets[:0]
}

//InitializeService initial service
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
}

//GetTweets return the tweets
func GetTweets() []*domain.Tweet {
	return tweets
}
