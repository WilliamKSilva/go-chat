package chat

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
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

func (messagesChannel *UpdateMessagesChannel) NewListener() {
    messagesChannel.mu.Lock()

    messagesChannel.Listeners++

    messagesChannel.mu.Unlock()
}

func (messagesChannel *UpdateMessagesChannel) NotifyListeners() {
    for i := 0; i < messagesChannel.Listeners; i++ {
        messagesChannel.Channel <- true 
    }
}

func (messagesChannel *UpdateMessagesChannel) Listening(messages []Message, conn *websocket.Conn) {
    msg := <-messagesChannel.Channel

    log.Println(msg)

    if msg {
        messagesResponse := ListMessagesResponse {
            Messages: messages,
        }

        data, err := json.Marshal(&messagesResponse)

        if err != nil {
            log.Println("Failed to notify client")
            return
        }

        conn.WriteMessage(websocket.BinaryMessage, data)
    }
}
