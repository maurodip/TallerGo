package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/TallerGo/src/domain"
)

type TweetManager struct {
	Tweets   []*domain.Tweet
	Users    map[string]*domain.User
	Topics   map[string]int
	Messages map[int]string
}

//PublishTweet publish a tweet
func (tm *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.Text == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	_, ok := tm.Users[tweet.User]

	if !ok {
		user := domain.NewUser(tweet.User)
		tm.Users[tweet.User] = user
	}

	listOfTweets := tm.Users[tweet.User].Tweets
	listOfTweets = append(listOfTweets, tweet)
	tm.Users[tweet.User].Tweets = listOfTweets

	tweet.Id = len(tm.Tweets) + 1
	tm.Tweets = append(tm.Tweets, tweet)

	trimmedTweet := strings.Fields(tweet.Text)
	for _, value := range trimmedTweet {
		count, ok := tm.Topics[value]
		if !ok {
			tm.Topics[value] = 0
		}
		tm.Topics[value] = count + 1
	}

	return tweet.Id, nil
}

//GetTweet return the tweet
func (tm TweetManager) GetTweet() *domain.Tweet {
	if len(tm.Tweets) != 0 {
		return tm.Tweets[len(tm.Tweets)-1]
	}
	return nil
}

//GetTweet return the tweet
func (tm TweetManager) GetTrendingTopics() []string {
	// var tt1, tt2 string
	// var
	// for key, value := range tm.Topics {

	// }
	// return nil
	n := map[int][]string{}
	var a []int
	for key, count := range tm.Topics {
		n[count] = append(n[count], key)
	}
	for k := range n {
		a = append(a, k)
	}

	var result []string
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			result = append(result, s)
		}
	}
	return result[:2]
}

//ClearTweet clear a tweet
func (tm TweetManager) ClearTweet() {
	tm.Tweets = tm.Tweets[:0]
	for _, user := range tm.Users {
		user.Tweets = make([]*domain.Tweet, 0)
	}
}

func NewTweetManager() *TweetManager {
	return &TweetManager{
		Tweets:   make([]*domain.Tweet, 0),
		Users:    make(map[string]*domain.User),
		Topics:   make(map[string]int),
		Messages: make(map[int]string),
	}
}

//GetTweets return the tweets
func (tm TweetManager) GetTweets() []*domain.Tweet {
	return tm.Tweets
}

//GetTweetById return the tweet by id
func (tm TweetManager) GetTweetById(id int) *domain.Tweet {
	return tm.Tweets[id-1]
}

func (tm TweetManager) CountTweetsByUser(user string) (int, error) {
	aUser, ok := tm.Users[user]
	if !ok {
		return 0, fmt.Errorf("user not exist")
	}
	return len(aUser.Tweets), nil
}

func (tm TweetManager) GetTweetsByUser(user string) ([]*domain.Tweet, error) {
	aUser, ok := tm.Users[user]
	if !ok {
		return nil, fmt.Errorf("user not exist")
	}
	return aUser.Tweets, nil
}

func (tm TweetManager) Follow(user string, userToFollow string) error {
	aUser, ok := tm.Users[user]
	if !ok {
		return fmt.Errorf("user not exist")
	}

	_, ok = tm.Users[userToFollow]
	if !ok {
		return fmt.Errorf("user not exist")
	}

	follows := aUser.Follows
	follows = append(follows, userToFollow)
	aUser.Follows = follows

	return nil
}

func (tm TweetManager) GetTimeline(user string) []*domain.Tweet {
	aUser, ok := tm.Users[user]
	listOfTweets := make([]*domain.Tweet, 0)
	if ok {
		// println("OK")
		for _, follows := range aUser.Follows {
			// println(follows)
			listOfTweets = append(listOfTweets, tm.Users[follows].Tweets...)
		}
		listOfTweets = append(listOfTweets, aUser.Tweets...)
	}
	return listOfTweets
}

func (tm *TweetManager) SendDirectMessage(user, userTo, msg string) {
	message := domain.NewMessage(user, msg)
	aUser := tm.Users[userTo]
	tm.Messages[message.Id] = userTo
	aUser.ReceiveDirectMessage(message)
}

func (tm TweetManager) GetUnreadedDirectMessages(user string) []*domain.Message {
	return tm.Users[user].GetUnreadedDirectMessages()
}

func (tm TweetManager) GetAllDirectMessages(user string) []*domain.Message {
	return tm.Users[user].GetAllDirectMessages()
}

func (tm TweetManager) ReadDirectMessage(idMsg int) {
	user := tm.Users[tm.Messages[idMsg]]
	user.GetMessage(idMsg).ReadMessage()
}
