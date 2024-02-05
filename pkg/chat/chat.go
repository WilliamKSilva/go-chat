package chat

import (
	"log"
	"slices"
	"sync"
)

type Message struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

type DeleteMessageRequest struct {
	MessageId string `json:"messageId"`
}

type ListMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type UpdateMessagesChannel struct {
    mu sync.Mutex
    Channel chan bool
    Listeners int
}

type Chat struct {
	Messages []Message
    UpdateMessagesChannel UpdateMessagesChannel
}

func (chat *Chat) NewMessage(message Message) {
	chat.Messages = append(chat.Messages, message)
}

func (chat *Chat) DeleteMessage(id string) {
	for i, message := range chat.Messages {
		if message.ID == id {
			chat.Messages = slices.Delete(chat.Messages, i, i+1)
		}
	}
}

func (chat *Chat) IsEmpty() bool {
	return len(chat.Messages) == 0
}

func (messagesChannel *UpdateMessagesChannel) NewListener() {
    messagesChannel.mu.Lock()

    messagesChannel.Listeners++

    messagesChannel.mu.Unlock()
}

func (messagesChannel *UpdateMessagesChannel) NotifyListeners() {
    for i := 0; i < messagesChannel.Listeners; i++ {
        log.Println("teste")
        messagesChannel.Channel <- true 
    }
}
