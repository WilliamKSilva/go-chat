package handler 

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	. "github.com/WilliamKSilva/go-chat/pkg/chat"
	"github.com/gorilla/websocket"
)

type HTMLFile struct {
    Data []byte
}

type HttpHandler struct {
    Chat Chat
    HtmlFile HTMLFile 
}

var upgrader = websocket.Upgrader{}
var internalServerError string = "Internal server error"

func (httpHandler *HttpHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete message handler!")
	var deleteMessageRequest DeleteMessageRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&deleteMessageRequest)

	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(internalServerError))
		return
	}

	httpHandler.Chat.DeleteMessage(deleteMessageRequest.MessageId)
}

func (httpHandler *HttpHandler) ListMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: This will be used by the Javascript script file to display the content
	// on the HTML page. Maybe there is an better way of injecting the data
	// on HTML directly from the server but I dont know yet.
	listMessagesResponse := ListMessagesResponse{
		Messages: httpHandler.Chat.Messages,
	}

	data, err := json.Marshal(listMessagesResponse)

	if err != nil {
		log.Println(err.Error())
		w.Write([]byte(internalServerError))
		return
	}

	w.Write(data)
}

func (httpHandler *HttpHandler) Html(w http.ResponseWriter, r *http.Request) {
	w.Write(httpHandler.HtmlFile.Data)
}

func (httpHandler *HttpHandler) Websocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

    log.Println("Client connected!")

	if err != nil {
		log.Println("Error upgrading websocket request")
		return
	}

	defer conn.Close()

    err = conn.WriteMessage(websocket.TextMessage, []byte("Hello client!"))

    if err != nil {
        log.Println("Error sending message")
        return
    }

    httpHandler.Chat.MessagesChannel.NewListener()

	for {
        // Listening to messages list update channel notify
        go httpHandler.Chat.MessagesChannel.Listening(httpHandler.Chat.Messages, conn)

		_, data, err := conn.ReadMessage()
        
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
            httpHandler.Chat.MessagesChannel.NotifyListeners()
            continue
		}

		if httpHandler.Chat.IsEmpty() {
			message.ID = "1"

			log.Println("recv: ", message.Nickname)
			log.Println("recv: ", message.Content)
			httpHandler.Chat.NewMessage(message)

			continue
		}

		lastMessage := httpHandler.Chat.Messages[len(httpHandler.Chat.Messages)-1]
		lastMessageId, err := strconv.Atoi(lastMessage.ID)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(internalServerError))
			break
		}

		message.ID = strconv.Itoa(lastMessageId + 1)

		httpHandler.Chat.NewMessage(message)
	}
}
