package main

import (
	"log"
	"github.com/gorilla/websocket"
	"flag"
	"fmt"
	"github.com/rcoh/terminal-watcher/types"
)


func readLoop(c *websocket.Conn) {
 for {
     if _, _, err := c.NextReader(); err != nil {
        c.Close()
          break
    }
 }
}

func sendMessage(ws *websocket.Conn, message types.Message) {
	text, err := message.Serialize()
	if err != nil {
		panic("Failed to serialize message")
	}
	println(string(text[:]))
	if err := ws.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
		log.Println("WriteMessage: %v", err)
    }
}


func main() {
	const message = "Hello World!"
	const server = "130.211.141.47"
	//const server = "127.0.0.1:8080"

	dialer := websocket.DefaultDialer;
	ws, _, err := dialer.Dial("ws://" + server + "/ws", nil)
	var mode = flag.String("mode", "", "modes (s: start, e: end)")
	var command = flag.String("command", "", "Command that was run (and finished)")
	var status = flag.Int("status", 0, "Status of command run")
	var clientId = flag.String("client", "", "ClientId")
	flag.Parse()
	defer ws.Close()

	go readLoop(ws)
	if err != nil {
		log.Println("Connection and sending failed", err);
	}

	fmt.Println(*mode);
	if *mode == "s" {
		sendMessage(ws, types.StartMessage(*command, *clientId))
	} else if *mode == "e" {
		sendMessage(ws, types.EndMessage(*command, *status, *clientId))
	}
}
