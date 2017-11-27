package domain

type User struct {
	User     string
	Follows  []string
	Tweets   []Tweet
	Messages map[int]*Message
	Favs     []Tweet
}

func NewUser(aUser string) *User {

	return &User{User: aUser, Tweets: make([]Tweet, 0), Messages: make(map[int]*Message), Favs: make([]Tweet, 0)}
}

func (user *User) ReceiveDirectMessage(message *Message) {
	user.Messages[message.Id] = message
}

func (user *User) GetUnreadedDirectMessages() []*Message {
	var unreadedMsg []*Message
	for _, msg := range user.Messages {
		if !msg.WasReaded {
			unreadedMsg = append(unreadedMsg, msg)
		}
	}
	return unreadedMsg
}

func (user *User) GetAllDirectMessages() []*Message {
	var unreadedMsg []*Message
	for _, msg := range user.Messages {
		unreadedMsg = append(unreadedMsg, msg)
	}
	return unreadedMsg
}

func (user *User) GetMessage(idMsg int) *Message {
	return user.Messages[idMsg]
}
