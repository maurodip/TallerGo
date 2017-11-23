package service

import (
	"fmt"

	"github.com/TallerGo/src/domain"
)

var tweets []*domain.Tweet
var tweetsByUser map[string][]*domain.Tweet

//PublishTweet publish a tweet
func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	_, ok := tweetsByUser[tweet.User]

	if !ok {
		tweetsByUser[tweet.User] = make([]*domain.Tweet, 0)
	}

	listByUser := tweetsByUser[tweet.User]
	listByUser = append(listByUser, tweet)
	tweetsByUser[tweet.User] = listByUser

	tweet.Id = len(tweets) + 1
	tweets = append(tweets, tweet)
	return tweet.Id, nil
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
	tweetsByUser = make(map[string][]*domain.Tweet)
}

//InitializeService initial service
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	tweetsByUser = make(map[string][]*domain.Tweet)
}

//GetTweets return the tweets
func GetTweets() []*domain.Tweet {
	return tweets
}

//GetTweetById return the tweet by id
func GetTweetById(id int) *domain.Tweet {
	return tweets[id-1]
}

func CountTweetsByUser(user string) (int, error) {
	list, ok := tweetsByUser[user]
	if !ok {
		return 0, fmt.Errorf("user not exist")
	}
	return len(list), nil
}

func GetTweetsByUser(user string) ([]*domain.Tweet, error) {
	list, ok := tweetsByUser[user]
	if !ok {
		return nil, fmt.Errorf("user not exist")
	}
	return list, nil
}
