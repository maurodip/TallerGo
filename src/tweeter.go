package main

import (
	"strconv"

	"github.com/TallerGo/src/domain"
	"github.com/TallerGo/src/rest"
	"github.com/TallerGo/src/service"
	"github.com/abiosoft/ishell"
)

func main() {

	fileTW := service.NewFileTweetWriter()
	channelTW := service.NewChannelTweetWriter(fileTW)
	tweetManager := service.NewTweetManager(channelTW)

	ginServer := rest.NewGinServer(tweetManager)
	ginServer.StartGinServer()

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your User: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			var tweet domain.Tweet
			tweet = domain.NewTextTweet(user, text)
			quit := make(chan bool)
			_, err := tweetManager.PublishTweet(tweet, quit)

			if err == nil {
				c.Print("Tweet sent\n")
			} else {
				c.Println(err.Error())
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweet()

			if tweet != nil {
				c.Println(tweet)
			} else {
				c.Println("Don't have tweet to show")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweets",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write User to count: ")

			user := c.ReadLine()
			count, _ := tweetManager.CountTweetsByUser(user)

			c.Println(count)
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "clearTweet",
		Help: "Clean a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			tweetManager.ClearTweet()

			c.Println("Tweet has been deleted")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "retweetTweet",
		Help: "Retweet a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write id tweet: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweetManager.Retweetear(id, user)

			c.Println("Tweet has been retweeted")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "follow",
		Help: "Follow a user",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			c.Print("Write user to follow: ")

			userToFollow := c.ReadLine()

			tweetManager.Follow(user, userToFollow)

			c.Println("Add followed user")
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "timeline",
		Help: "Get a user timeline",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your user: ")

			user := c.ReadLine()

			timelineTweets := tweetManager.GetTimeline(user)

			for _, v := range timelineTweets {
				c.Println(v.PrintableTweet())
			}
			return
		},
	})

	shell.Run()
}
