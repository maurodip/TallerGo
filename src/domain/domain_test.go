package domain_test

import (
	"testing"

	"github.com/TallerGo/src/domain"
)

func TestCanGetAPrinteableTweet(t *testing.T) {
	//Initialization
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	//Operation
	text := tweet.PrintableTweet()

	//Validation
	expectedText := "@grupoesfera: This is my tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s", expectedText, text)
	}
}
