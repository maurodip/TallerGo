package service

import "github.com/TallerGo/src/domain"

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{Tweets: make([]domain.Tweet, 0)}
}

func (memoryTW *MemoryTweetWriter) WriteTweet(tweet domain.Tweet) {
	memoryTW.Tweets = append(memoryTW.Tweets, tweet)
}
