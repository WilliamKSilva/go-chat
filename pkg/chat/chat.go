package chat

import (
	"log"
	"slices"
	"sync"
)

type Chat struct {
    mu sync.Mutex
	Messages []Message
    MessagesChannel MessagesChannel
}

func (chat *Chat) NewMessage(message Message) {
    chat.mu.Lock()
	chat.Messages = append(chat.Messages, message)
    log.Println("New message added!")
    chat.mu.Unlock()
}

func (chat *Chat) DeleteMessage(id string) {
    chat.mu.Lock()
	for i, message := range chat.Messages {
		if message.ID == id {
			chat.Messages = slices.Delete(chat.Messages, i, i+1)
		}
	}
    chat.mu.Unlock()
}

func (chat *Chat) IsEmpty() bool {
    chat.mu.Lock()
    isEmpty := len(chat.Messages) == 0
    chat.mu.Unlock()

    return isEmpty
}


