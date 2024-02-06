package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/url"

	. "github.com/WilliamKSilva/go-chat/pkg/chat"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address") 

func sendMessage(c *websocket.Conn, data []byte) {
    err := c.WriteMessage(websocket.TextMessage, data)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func connectWebsocket() {
    u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
    c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer c.Close()

    message := Message{
        Nickname: "ddos",
        Content: "This is an spam test",
    }
    data, err := json.Marshal(message)
    if err != nil {
        log.Println(err.Error())
    }

    for {
        sendMessage(c, data)
        
        log.Println("Message sent!")
    }
}

func main() {
    for {
        go connectWebsocket()
    }
}
