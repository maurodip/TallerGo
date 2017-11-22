package main

import (
	"github.com/TallerGo/src/domain"
	"github.com/TallerGo/src/service"
	"github.com/abiosoft/ishell"
)

func main() {

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

			var tweet *domain.Tweet
			tweet = domain.NewTweet(user, text)

			err := service.PublishTweet(tweet)

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

			tweet := service.GetTweet()

			if tweet != nil {
				c.Println(tweet.Text)
			} else {
				c.Println("Don't have tweet to show")
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "clearTweet",
		Help: "Clean a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			service.ClearTweet()

			c.Println("Tweet has been deleted")
			return
		},
	})
	shell.Run()
}
