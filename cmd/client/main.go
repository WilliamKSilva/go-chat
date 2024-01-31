package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"

	chat "github.com/WilliamKSilva/go-chat/internal"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address") 

func main() {
    u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal("Failed to connect to chat")
    }
    defer c.Close()

    log.Fatal("Connected to chat!")

    for {
        message := chat.Message{
            Nickname: "ddos",
            Content: "This is an spam test",
        }
        data, err := json.Marshal(message)
        if err != nil {
            log.Println("Error on data Marshal")
        }

        log.Println(data)

        err = c.WriteMessage(websocket.TextMessage, data)
        if err != nil {
            log.Println("Error trying to send chat message")
        }

        log.Println("Message sent!")
    }
    
}
