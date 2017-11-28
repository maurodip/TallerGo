package service

import "github.com/TallerGo/src/domain"

type ChannelTweetWriter struct {
	TweetWriter TweetWriter
}

func NewChannelTweetWriter(tw TweetWriter) *ChannelTweetWriter {
	return &ChannelTweetWriter{TweetWriter: tw}
}

func (channel *ChannelTweetWriter) WriteTweet(tweetsToWrite chan domain.Tweet, quit chan bool) {
	tweet, open := <-tweetsToWrite
	for open {
		channel.TweetWriter.WriteTweet(tweet)
		tweet, open = <-tweetsToWrite
	}
	quit <- true
}
