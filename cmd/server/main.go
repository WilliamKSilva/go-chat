package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gorilla/websocket"
    . "github.com/WilliamKSilva/go-chat/pkg/chat"
)

type DeleteMessageRequest struct {
	MessageId string `json:"messageId"`
}

type ListMessagesResponse struct {
	Messages []Message `json:"messages"`
}

type HTMLFile struct {
    data []byte
}

var chat Chat

var upgrader = websocket.Upgrader{}

var internalServerError = "Internal server error"

func connectChat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading websocket request")
		return
	}

	defer c.Close()

    err = c.WriteMessage(websocket.TextMessage, []byte("Hello client!"))

    if err != nil {
        log.Println("Error sending message")
        return
    }

	for {
		_, data, err := c.ReadMessage()
        
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(internalServerError))
			break
		}

		var message Message 
		err = json.Unmarshal(data, &message)

        // If there is an error the content is probably plain text
		if err != nil {
            log.Println(string(data))
            return
		}

		if chat.IsEmpty() {
			message.ID = "1"

			log.Println("recv: ", message.Nickname)
			log.Println("recv: ", message.Content)
			chat.NewMessage(message)

			continue
		}

		lastMessage := chat.Messages[len(chat.Messages)-1]
		lastMessageId, err := strconv.Atoi(lastMessage.ID)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(internalServerError))
			break
		}

		message.ID = strconv.Itoa(lastMessageId + 1)

		log.Println("recv: ", message.Nickname)
		log.Println("recv: ", message.Content)
		chat.NewMessage(message)

	    log.Println(len(chat.Messages))
	}
}

func deleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete message handler!")
	var deleteMessageRequest DeleteMessageRequest

    decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&deleteMessageRequest)

	if err != nil {
        log.Println(err.Error())
		w.Write([]byte(internalServerError))
		return
	}

	chat.DeleteMessage(deleteMessageRequest.MessageId)
}

func listMessagesHandler(w http.ResponseWriter, r *http.Request) {
   // TODO: This will be used by the Javascript script file to display the content
   // on the HTML page. Maybe there is an better way of injecting the data
   // on HTML directly from the server but I dont know yet.
   listMessagesResponse := ListMessagesResponse{
       Messages: chat.Messages,
   }

   data, err := json.Marshal(listMessagesResponse)

   if err != nil {
       log.Println(err.Error())
       w.Write([]byte(internalServerError))
       return
   }

   w.Write(data)
}

func (f HTMLFile) chatHandler(w http.ResponseWriter, r *http.Request) {
    w.Write(f.data)
}

func main() {
    gp := os.Getenv("GOPATH")
    htmlPath := path.Join(gp, "web/index.html")

    data, err := os.ReadFile(htmlPath)

    if err != nil {
        log.Fatal(err.Error())
    }

    htmlFile := HTMLFile{
        data: data,
    }

    if err != nil {
        log.Fatal(err.Error())
    }

    message := Message{
        Nickname: "test",
        Content: "teste",
    }
    chat.NewMessage(message)

	http.HandleFunc("/chat", connectChat)
	http.HandleFunc("/delete-message", deleteMessageHandler)
	http.HandleFunc("/list-message", listMessagesHandler)
	http.HandleFunc("/", htmlFile.chatHandler)
	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
