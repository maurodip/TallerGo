package domain

import (
	"time"
)

var idAutonum int

type Message struct {
	Id        int
	User      string
	Text      string
	WasReaded bool
	Date      *time.Time
}

func NewMessage(userFrom string, text string) *Message {

	date := time.Now()
	idAutonum++
	return &Message{Id: idAutonum, User: userFrom, Text: text, Date: &date, WasReaded: false}
}

func (m *Message) ReadMessage() {
	m.WasReaded = true
}
