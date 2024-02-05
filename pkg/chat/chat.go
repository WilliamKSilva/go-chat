package chat

import "slices"

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

type Chat struct {
	Messages []Message
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
