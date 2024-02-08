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
	Data []Message `json:"messages"`
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
    mc.mu.Lock()
    for i := 0; i < mc.Listeners; i++ {
        mc.Channel <- true 
    }
    mc.mu.Unlock()
}

func (mc *MessagesChannel) Listening(messages *[]Message, conn *websocket.Conn) {
    for {
        msg := <-mc.Channel

        if msg {
            log.Println("Updating client list")
            messagesResponse := ListMessagesResponse {
                Data: *messages,
            }

            data, err := json.Marshal(&messagesResponse)

            if err != nil {
                log.Println("Failed to notify client")
                return
            }

            conn.WriteMessage(websocket.BinaryMessage, data)
        }
    }
}
