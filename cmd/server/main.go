package main

import (
	"log"
	"net/http"
	"os"
	"path"

    . "github.com/WilliamKSilva/go-chat/pkg/chat"
    . "github.com/WilliamKSilva/go-chat/pkg/handler"
)

var httpHandler HttpHandler

func main() {
    httpHandler.Chat = Chat{} 
    httpHandler.Chat.MessagesChannel = MessagesChannel{
        Channel: make(chan bool),
        Listeners: 0,
    }

    gp := os.Getenv("GOPATH")
    htmlPath := path.Join(gp, "web/index.html")

    data, err := os.ReadFile(htmlPath)

    if err != nil {
        log.Fatal(err.Error())
    }

    htmlFile := HTMLFile{
        Data: data,
    }

    httpHandler.HtmlFile = htmlFile

    if err != nil {
        log.Fatal(err.Error())
    }

	http.HandleFunc("/chat", httpHandler.Websocket)
	http.HandleFunc("/delete-message", httpHandler.DeleteMessage)
	http.HandleFunc("/list-message", httpHandler.ListMessages)
	http.HandleFunc("/", httpHandler.Html)
	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
