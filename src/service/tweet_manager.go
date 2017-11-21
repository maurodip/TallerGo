package service

var tweet string

func PublishTweet(aTweet string) {
	tweet = aTweet
}

func GetTweet() string {
	return tweet
}
