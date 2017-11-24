package domain

type Tweet interface {
	GetUser() string
	GetText() string
	GetId() int
	Retweet()
	Favear()
	PrintableTweet() string
	String()
}
