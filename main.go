package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func connectWS(w http.ResponseWriter, r *http.Request) {
    c, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Error upgrading websocket request")
        return
    }
    
    defer c.Close()
    for {
        _, message, err := c.ReadMessage()

        if err != nil {
            log.Println("Error reading message")
            response := "Internal server error"
            w.Write([]byte(response))
            break
        }

        log.Println("recv: ", string(message))
    }
}

func main() {
    http.HandleFunc("/", connectWS)
    log.Println("Server listening on port: 8080")
    http.ListenAndServe(":8080", nil)
}
