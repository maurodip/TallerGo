package service

import (
	"fmt"

	"github.com/TallerGo/src/domain"
)

var tweets []*domain.Tweet

//var tweetsByUser map[string][]*domain.Tweet
var users map[string]*domain.User

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

	_, ok := users[tweet.User]

	if !ok {
		//tweetsByUser[tweet.User] = make([]*domain.Tweet, 0)
		user := domain.NewUser(tweet.User)
		users[tweet.User] = user
	}

	// listByUser := tweetsByUser[tweet.User]
	// listByUser = append(listByUser, tweet)
	// tweetsByUser[tweet.User] = listByUser

	listOfTweets := users[tweet.User].Tweets
	listOfTweets = append(listOfTweets, tweet)
	users[tweet.User].Tweets = listOfTweets

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
	//tweetsByUser = make(map[string][]*domain.Tweet)
	for _, user := range users {
		user.Tweets = make([]*domain.Tweet, 0)
	}
}

//InitializeService initial service
func InitializeService() {
	tweets = make([]*domain.Tweet, 0)
	//tweetsByUser = make(map[string][]*domain.Tweet)
	users = make(map[string]*domain.User)
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
	aUser, ok := users[user]
	if !ok {
		return 0, fmt.Errorf("user not exist")
	}
	return len(aUser.Tweets), nil
}

func GetTweetsByUser(user string) ([]*domain.Tweet, error) {
	aUser, ok := users[user]
	if !ok {
		return nil, fmt.Errorf("user not exist")
	}
	return aUser.Tweets, nil
}

func Follow(user string, userToFollow string) error {
	aUser, ok := users[user]
	if !ok {
		return fmt.Errorf("user not exist")
	}

	_, ok = users[userToFollow]
	if !ok {
		return fmt.Errorf("user not exist")
	}

	follows := aUser.Follows
	follows = append(follows, userToFollow)
	aUser.Follows = follows

	return nil
}

func GetTimeline(user string) []*domain.Tweet {
	aUser, ok := users[user]
	listOfTweets := make([]*domain.Tweet, 0)
	if ok {
		// println("OK")
		for _, follows := range aUser.Follows {
			// println(follows)
			listOfTweets = append(listOfTweets, users[follows].Tweets...)
		}
		listOfTweets = append(listOfTweets, aUser.Tweets...)
	}
	return listOfTweets
}
