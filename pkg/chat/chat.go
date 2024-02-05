package chat

import (
	"slices"
)

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


