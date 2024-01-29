package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
    Nickname string `json:"nickname"`
    Content string `json:"content"`
}

type Chat struct {
    messages []Message
}

var chat Chat

var upgrader = websocket.Upgrader{}

var internalServerError = "Internal server error"

func (chat *Chat) newMessage(message Message) {
    chat.messages = append(chat.messages, message)
}

func connectWS(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Error upgrading websocket request")
        return
    }

    defer c.Close()
    for {
        _, data, err := c.ReadMessage()

        if err != nil {
            log.Println("Error reading message")
            w.Write([]byte(internalServerError))
            break
        }

        var message Message
        err = json.Unmarshal(data, &message)
        if err != nil {
            log.Println("Unprocessable data")
            w.Write([]byte(internalServerError))
            break
        }

        log.Println("recv: ", message.Nickname)
        log.Println("recv: ", message.Content)

        chat.newMessage(message)

        log.Println(len(chat.messages))
    }
}

func main() {
    http.HandleFunc("/", connectWS)
    log.Println("Server listening on port: 8080")
    http.ListenAndServe(":8080", nil)
}
