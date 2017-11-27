package domain

import (
	"fmt"
)

type FacebookPlugin struct {
	Name string
}

func NewFacebookPlugin() *FacebookPlugin {
	return &FacebookPlugin{Name: "Facebook"}
}

func (facebookPlugin *FacebookPlugin) Share(tweet Tweet) {
	fmt.Println("Se compartio en facebook: " + tweet.String())
}
