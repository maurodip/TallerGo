package service

import (
	"fmt"
	"sort"
	"strings"

	"github.com/TallerGo/src/domain"
)

type TweetManager struct {
	TweetWriter *ChannelTweetWriter
	Tweets      []domain.Tweet
	Users       map[string]*domain.User
	Topics      map[string]int
	Messages    map[int]string
	Plugin      []domain.TweetPlugin
}

func NewTweetManager(channelTW *ChannelTweetWriter) *TweetManager {
	return &TweetManager{
		TweetWriter: channelTW,
		Tweets:      make([]domain.Tweet, 0),
		Users:       make(map[string]*domain.User),
		Topics:      make(map[string]int),
		Messages:    make(map[int]string),
		Plugin:      make([]domain.TweetPlugin, 0),
	}
}

//PublishTweet publish a tweet
func (tm *TweetManager) PublishTweet(tweet domain.Tweet, quit chan bool) (int, error) {
	if tweet.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}
	if tweet.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	}

	_, ok := tm.Users[tweet.GetUser()]

	if !ok {
		user := domain.NewUser(tweet.GetUser())
		tm.Users[tweet.GetUser()] = user
	}

	listOfTweets := tm.Users[tweet.GetUser()].Tweets
	listOfTweets = append(listOfTweets, tweet)
	tm.Users[tweet.GetUser()].Tweets = listOfTweets

	tweet.SetId(len(tm.Tweets) + 1)
	tm.Tweets = append(tm.Tweets, tweet)
	tm.ApplyPlugins(tweet)
	trimmedTweet := strings.Fields(tweet.GetText())
	for _, value := range trimmedTweet {
		count, ok := tm.Topics[value]
		if !ok {
			tm.Topics[value] = 0
		}
		tm.Topics[value] = count + 1
	}

	tweetsToWrite := make(chan domain.Tweet)
	go tm.TweetWriter.WriteTweet(tweetsToWrite, quit)

	tweetsToWrite <- tweet
	close(tweetsToWrite)

	return tweet.GetId(), nil
}

//GetTweet return the tweet
func (tm TweetManager) GetTweet() domain.Tweet {
	if len(tm.Tweets) != 0 {
		return tm.Tweets[len(tm.Tweets)-1]
	}
	return nil
}

//GetTweet return the tweet
func (tm TweetManager) GetTrendingTopics() []string {
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
		user.Tweets = make([]domain.Tweet, 0)
	}
}

//GetTweets return the tweets
func (tm TweetManager) GetTweets() []domain.Tweet {
	return tm.Tweets
}

//GetTweetById return the tweet by id
func (tm TweetManager) GetTweetById(id int) domain.Tweet {
	return tm.Tweets[id-1]
}

func (tm TweetManager) CountTweetsByUser(user string) (int, error) {
	aUser, ok := tm.Users[user]
	if !ok {
		return 0, fmt.Errorf("user not exist")
	}
	return len(aUser.Tweets), nil
}

func (tm TweetManager) GetTweetsByUser(user string) ([]domain.Tweet, error) {
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

func (tm TweetManager) GetTimeline(user string) []domain.Tweet {
	aUser, ok := tm.Users[user]
	listOfTweets := make([]domain.Tweet, 0)
	if ok {
		for _, follows := range aUser.Follows {
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

func (tm TweetManager) Retweetear(idtweet int, user string) {
	var tweet domain.Tweet
	tweet = tm.GetTweetById(idtweet)
	tweet.Retweet()
	aUser := tm.Users[user]
	aUser.Tweets = append(aUser.Tweets, tweet)
}

func (tm TweetManager) Favear(idtweet int, user string) {
	tweet := tm.GetTweetById(idtweet)
	tweet.Favear()
	aUser := tm.Users[user]
	aUser.Favs = append(aUser.Favs, tweet)
}

func (tm TweetManager) GetTweetsFavs(user string) []domain.Tweet {
	return tm.Users[user].Favs
}

func (tm *TweetManager) SetPlugin(plugin domain.TweetPlugin) {
	tm.Plugin = append(tm.Plugin, plugin)
}

func (tm TweetManager) ApplyPlugins(tweet domain.Tweet) {
	for _, p := range tm.Plugin {
		p.Share(tweet)
	}
}
