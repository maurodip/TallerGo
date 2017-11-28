package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	Id       int
	User     string
	Text     string
	Date     *time.Time
	Retweets int
	Favs     int
}

func NewTextTweet(aUser string, text string) *TextTweet {

	date := time.Now()

	return &TextTweet{User: aUser, Text: text, Date: &date}
}

func (tweet *TextTweet) Retweet() {
	tweet.Retweets++
}

func (tweet *TextTweet) Favear() {
	tweet.Favs++
}

func (tweet *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("id:%d @%s: %s", tweet.Id, tweet.User, tweet.Text)
}

func (tweet *TextTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}
