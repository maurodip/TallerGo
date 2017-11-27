package domain

import "time"

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(id int)
	Retweet()
	Favear()
	PrintableTweet() string
	String() string
}
