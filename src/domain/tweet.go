package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	Id       int
	User     string
	Text     string
	Date     *time.Time
	Retweets int
	Favs     int
}

func NewTweet(aUser string, text string) *Tweet {

	date := time.Now()

	return &Tweet{User: aUser, Text: text, Date: &date}
}

func (tweet Tweet) Retweet() {
	tweet.Retweets++
}

func (tweet Tweet) Favear() {
	tweet.Favs++
}

func (tweet Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", tweet.User, tweet.Text)
}
