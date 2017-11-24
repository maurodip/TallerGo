package domain

import (
	"fmt"
	"time"
)

type QuoteTweet struct {
	Id       int
	User     string
	Text     string
	Date     *time.Time
	Retweets int
	Favs     int
	Quote    *TextTweet
}

func NewQuoteTweet(aUser string, text string, quote *TextTweet) *QuoteTweet {

	date := time.Now()

	return &QuoteTweet{User: aUser, Text: text, Date: &date, Quote: quote}
}

func (tweet *QuoteTweet) Retweet() {
	tweet.Retweets++
}

func (tweet *QuoteTweet) Favear() {
	tweet.Favs++
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`@%s: %s "@%s: %s"`, tweet.User, tweet.Text, tweet.Quote.User, tweet.Quote.Text)
} //`@nick: Awesome "@grupoesfera: This is my tweet"`

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}
