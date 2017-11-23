package service_test

import (
	"testing"

	"github.com/TallerGo/src/domain"
	"github.com/TallerGo/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()
	println(publishedTweet == nil)
	isValidTweet(t, publishedTweet, id, user, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members
	   all over the world. To date all community oriented activities have been organized by the community
	   with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet *domain.Tweet

	user := "grupoesfera"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count, _ := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	firstId, _ := tweetManager.PublishTweet(tweet)
	secondId, _ := tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets, _ := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, user, secondText) {
		return
	}

}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

	if tweet.User != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}

func TestCannotRetrieveTheTweetsSentByAnUserNotExisting(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	user := "grupoesfera"

	// Operation
	_, err := tweetManager.GetTweetsByUser(user)

	// Validation
	if err == nil && err.Error() == "user not exist" {

		t.Errorf("Error expected is user not exist")
		return
	}

}

func TestCannotCountTheTweetsSentByAnUserNotExist(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet

	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	_, err := tweetManager.CountTweetsByUser("montoto")

	// Validation
	if err == nil && err.Error() == "user not exist" {
		t.Errorf("Error expected is user not exist")
		return
	}

}

func TestUserCanFollowUsers(t *testing.T) {
	tm := service.NewTweetManager()
	user := "nportas"
	anotherUser := "mercadolibre"
	text := "This is my first tweet portas"
	secondText := "This is my second tweet meli"

	tweet := domain.NewTweet(user, text)
	secondTweet := domain.NewTweet(anotherUser, secondText)

	firstId, _ := tm.PublishTweet(tweet)
	secondId, _ := tm.PublishTweet(secondTweet)

	grupoesferaTweet := domain.NewTweet("grupoesfera", text)
	thirdId, _ := tm.PublishTweet(grupoesferaTweet)

	tm.Follow("grupoesfera", user)
	tm.Follow("grupoesfera", anotherUser)

	timeline := tm.GetTimeline("grupoesfera")

	println(timeline)
	if len(timeline) != 3 {
		t.Errorf("Error expected is user not exist")
		return
	}

	firstPublishedTweet := timeline[0]
	secondPublishedTweet := timeline[1]
	thirdPublishedTweet := timeline[2]

	if !isValidTweet(t, firstPublishedTweet, firstId, user, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, anotherUser, secondText) {
		return
	}

	if !isValidTweet(t, thirdPublishedTweet, thirdId, "grupoesfera", text) {
		return
	}
}
