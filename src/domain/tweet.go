package domain

import (
	"time"
)

type Tweet struct {
	Id   int
	User string
	Text string
	Date *time.Time
}

func NewTweet(aUser string, text string) *Tweet {

	date := time.Now()

	return &Tweet{User: aUser, Text: text, Date: &date}
}
