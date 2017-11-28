package service

import (
	"os"

	"github.com/TallerGo/src/domain"
)

type FileTweetWriter struct {
	file *os.File
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)

	writer := new(FileTweetWriter)
	writer.file = file

	return writer
}

func (fileTW *FileTweetWriter) WriteTweet(tweet domain.Tweet) {
	if fileTW.file != nil {
		byteSlice := []byte(tweet.PrintableTweet() + "\n")
		fileTW.file.Write(byteSlice)
	}
}
