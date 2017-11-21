package service

var tweet string

//PublishTweet publish a tweet
func PublishTweet(aTweet string) {
	tweet = aTweet
}

//GetTweet return the tweet
func GetTweet() string {
	return tweet
}

//ClearTweet clear tweet
func ClearTweet() {
	tweet = ""
}
