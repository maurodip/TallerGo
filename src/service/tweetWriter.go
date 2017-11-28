package service

import "github.com/TallerGo/src/domain"

type TweetWriter interface {
	WriteTweet(tweet domain.Tweet)
}
