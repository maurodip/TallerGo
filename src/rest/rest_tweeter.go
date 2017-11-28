package rest

import (
	"net/http"

	"github.com/TallerGo/src/domain"

	"github.com/TallerGo/src/service"
	"github.com/gin-gonic/gin"
)

type GinTweet struct {
	User string
	Text string
}

type GinServer struct {
	tweetManager *service.TweetManager
}

func NewGinServer(tweetManager *service.TweetManager) *GinServer {
	return &GinServer{tweetManager}
}

func (server *GinServer) StartGinServer() {

	router := gin.Default()

	router.GET("/listTweets", server.listTweets)
	router.GET("/listTweets/:user", server.getTweetsByUser)
	router.POST("publishTweet", server.publishTweet)

	go router.Run()
}

func (server *GinServer) listTweets(c *gin.Context) {

	c.JSON(http.StatusOK, server.tweetManager.GetTweets())
}

func (server *GinServer) getTweetsByUser(c *gin.Context) {

	user := c.Param("user")
	tweets, _ := server.tweetManager.GetTweetsByUser(user)
	c.JSON(http.StatusOK, tweets)
}

func (server *GinServer) publishTweet(c *gin.Context) {

	quit := make(chan bool)

	var tweetdata GinTweet
	c.Bind(&tweetdata)

	tweetToPublish := domain.NewTextTweet(tweetdata.User, tweetdata.Text)

	id, err := server.tweetManager.PublishTweet(tweetToPublish, quit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error publishing tweet "+err.Error())
	} else {
		c.JSON(http.StatusOK, struct{ Id int }{id})
	}
}
