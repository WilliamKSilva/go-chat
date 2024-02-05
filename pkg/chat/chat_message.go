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

type MessagesChannel struct {
    mu sync.Mutex
    Channel chan bool
    Listeners int
}

func (mc *MessagesChannel) NewListener() {
    mc.mu.Lock()

    mc.Listeners++

    mc.mu.Unlock()
}

func (mc *MessagesChannel) NotifyListeners() {
    for i := 0; i < mc.Listeners; i++ {
        mc.Channel <- true 
    }
}

func (mc*MessagesChannel) Listening(messages []Message, conn *websocket.Conn) {
    msg := <-mc.Channel

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
