package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	Id       int
	User     string
	Text     string
	Date     *time.Time
	Retweets int
	Favs     int
	Image    string
}

func NewImageTweet(aUser string, text string, image string) *ImageTweet {

	date := time.Now()

	return &ImageTweet{User: aUser, Text: text, Date: &date, Image: image}
}

func (tweet *ImageTweet) Retweet() {
	tweet.Retweets++
}

func (tweet *ImageTweet) Favear() {
	tweet.Favs++
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.Image)
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}
