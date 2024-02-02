# How to run

- You need Golang compiler/toolchain: https://go.dev/doc/install
- Launch the server: `go run cmd/server/main.go`. Optionally you can launch an websocket client that spawn multiple connections on the server to test if everything is ok: `go run cmd/client/main.go`.
The websocket client uses goroutines to spawn multiple concurrent tasks, so we can test multiple connections and send messages.
- You also can use the bash script `bash hotreload.sh` so everytime an file tracked by git is changed Golang compiler is called automatically.

# More
- The next step is to make an web page so people can actually send messages on an UI chat.
- I want to check another way of testing multiple websocket connections, at the moment its literally multiple goroutines being spawned from an while loop, I dont know if this is the better way of testing if the server will break.

![Captura de tela de 2024-01-31 11-13-44](https://github.com/WilliamKSilva/go-chat/assets/75429175/dfee9fbb-6ffd-43ca-88a6-d2b3f1f0dad0)
