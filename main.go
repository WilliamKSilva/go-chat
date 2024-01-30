package main

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/websocket"
)

type Message struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}

type Chat struct {
	messages []Message
}

type DeleteMessageRequest struct {
	MessageId string `json:"messageId"`
}

type ListMessagesResponse struct {
	Messages []Message `json:"messages"`
}

var chat Chat

var upgrader = websocket.Upgrader{}

var internalServerError = "Internal server error"

func (chat *Chat) newMessage(message Message) {
	chat.messages = append(chat.messages, message)
}

func (chat *Chat) deleteMessage(id string) {
	for i, message := range chat.messages {
		if message.ID == id {
            log.Println("Found!")
			chat.messages = slices.Delete(chat.messages, i, i+1)
		}
	}

	log.Println(len(chat.messages))
}

func (chat *Chat) isEmpty() bool {
	return len(chat.messages) == 0
}

func connectChat(w http.ResponseWriter, r *http.Request) {
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

		if chat.isEmpty() {
			log.Println("Chat is empty")
			message.ID = "1"

			log.Println("recv: ", message.Nickname)
			log.Println("recv: ", message.Content)
			chat.newMessage(message)
			log.Println(len(chat.messages))

			continue
		}

		lastMessage := chat.messages[len(chat.messages)-1]
		lastMessageId, err := strconv.Atoi(lastMessage.ID)

        log.Println(lastMessageId)

		if err != nil {
			log.Println("Error converting ID to number")
			w.Write([]byte(internalServerError))
			break
		}

		message.ID = strconv.Itoa(lastMessageId + 1)

		log.Println("recv: ", message.Nickname)
		log.Println("recv: ", message.Content)
		chat.newMessage(message)
		log.Println(len(chat.messages))
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

	chat.deleteMessage(deleteMessageRequest.MessageId)
}

func listMessagesHandler(w http.ResponseWriter, r *http.Request) {
   encoder := json.NewEncoder(w)

   listMessagesResponse := ListMessagesResponse{
       Messages: chat.messages,
   }

   err := encoder.Encode(&listMessagesResponse)

   if err != nil {
       log.Println(err.Error())
       w.Write([]byte(internalServerError))
       return
   }
}

func main() {
	http.HandleFunc("/", connectChat)
	http.HandleFunc("/delete-message", deleteMessageHandler)
	http.HandleFunc("/list-message", listMessagesHandler)
	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}
