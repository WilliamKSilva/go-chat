# How to run

- You need Golang compiler/toolchain: https://go.dev/doc/install
- Launch the server: ==go run cmd/server/main.go==. Optionally you can launch an websocket client that spawn multiple connections on the server to test if everything is ok: ==go run cmd/client/main.go==.
The websocket client uses goroutines to spawn multiple concurrent tasks, so we can test multiple connections and send messages.

# More
- The next step is to make an web page so people can actually send messages on an UI chat.
- I want to check another way of testing multiple websocket connections, at the moment its literally multiple goroutines being spawned from an while loop, I dont know if this is the better way of testing if the server will break.
