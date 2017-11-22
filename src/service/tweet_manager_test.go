package service_test

import (
	"testing"

	"github.com/TallerGo/src/domain"
	"github.com/TallerGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)

	publishedTweet := service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}
}
func TestCleanedTweetIsClean(t *testing.T) {

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	service.PublishTweet(tweet)
	service.ClearTweet()

	if service.GetTweet() != nil {
		t.Error("Expected empty tweet")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet

	var user string
	user = "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "grupoesfera"

	text := `Whether you use Waterfall, Agile, or Conversational Development, GitLab streamlines your collaborative workflows. 
	Visualize, prioritize, coordinate, and track your progress your way with GitLab’s flexible project management tools. 
	Spend less time configuring your tools, and more time creating. Whether you’re deploying to one server or thousands, 
	build, test, and release your code confidently and securely with GitLab’s built-in continuous delivery and deployment.`

	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil {
		t.Error("Expected error")
	}
	if err.Error() != "text exceeding 140 characters" {
		t.Error("Expected error is text exceeding 140 characters")
	}
}
