package domain

type TweetPlugin interface {
	Share(tweet Tweet)
}
