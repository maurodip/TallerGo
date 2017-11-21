package service_test

import (
	"testing"

	"github.com/TallerGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	tweet := "This is my fisrt tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is: ", tweet)
	}
}
