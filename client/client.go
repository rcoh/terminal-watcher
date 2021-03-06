package main

import (
	"log"
	"github.com/gorilla/websocket"
	"flag"
	"github.com/rcoh/terminal-watcher/types"
	"os"
	"bufio"
	"fmt"
	"github.com/rcoh/terminal-watcher/ratelimit"
)


func readLoop(c *websocket.Conn) {
 for {
     if _, _, err := c.NextReader(); err != nil {
        fmt.Println(err)
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

	if err := ws.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
		log.Println("WriteMessage: %v", err)
    }
}

func follow(ws *websocket.Conn, clientId string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rate_per_sec := 50
	limiter, _ := ratelimit.NewRateLimiter(rate_per_sec)
	for scanner.Scan() {
		// TODO: drop messages instead of just rate limiting
		text := scanner.Text()

		if limiter.Limit() {
			// log.Println("rate limiting")
		} else {
			message := types.TailMessage(text, clientId)
			sendMessage(ws, message)
		}
	}
}


func main() {
	const server = "130.211.141.47"
	//const server = "127.0.0.1:8080"

	dialer := websocket.DefaultDialer;
	ws, _, err := dialer.Dial("ws://" + server + "/ws", nil)
	var mode = flag.String("mode", "", "modes (s: start, e: end)")
	var command = flag.String("command", "", "Command that was run (and finished)")
	var status = flag.Int("status", 0, "Status of command run")
	var clientId = flag.String("client", "", "ClientId")
	var file = flag.String("file", "/tmp/follow", "Script output to follow")
	flag.Parse()
	defer ws.Close()

	go readLoop(ws)
	if err != nil {
		log.Println("Connection and sending failed", err);
	}

	if *mode == "s" {
		sendMessage(ws, types.StartMessage(*command, *clientId))
	} else if *mode == "e" {
		sendMessage(ws, types.EndMessage(*command, *status, *clientId))
	} else if *mode == "f" {
		follow(ws, *clientId, *file)
	}
}
