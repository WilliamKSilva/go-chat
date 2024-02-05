package main

import (
	"log"
	"net/http"
	"os"
	"path"

    . "github.com/WilliamKSilva/go-chat/pkg/chat"
    . "github.com/WilliamKSilva/go-chat/pkg/httpHandler"
)

var httpHandler HttpHandler

func main() {
    httpHandler.Chat = Chat{} 

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

    message := Message{
        Nickname: "test",
        Content: "teste",
    }

    httpHandler.Chat.NewMessage(message)

	http.HandleFunc("/chat", httpHandler.Websocket)
	http.HandleFunc("/delete-message", httpHandler.DeleteMessage)
	http.HandleFunc("/list-message", httpHandler.ListMessages)
	http.HandleFunc("/", httpHandler.Html)
	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
