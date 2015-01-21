package main

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/ActiveState/tail"
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

func tailOutput(ws *websocket.Conn) {
	t, _ := tail.TailFile("test", tail.Config{Follow: true})
	for line := range t.Lines {
	    if err := ws.WriteMessage(websocket.TextMessage, []byte(line.Text)); err != nil {
			log.Println("WriteMessage: %v", err)
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


func main() {
	const message = "Hello World!"
	const server = "130.211.141.47"
	dialer := websocket.DefaultDialer;
	ws, _, err := dialer.Dial("ws://" + server + "/ws", nil)
	var mode = flag.String("mode", "", "modes (-s: start, -e: end)")
	flag.Parse()
	defer ws.Close()

	go readLoop(ws)
	if err != nil {
		log.Println("Connection and sending failed", err);
	}

	fmt.Println(*mode);
	if *mode == "s" {
		sendMessage(ws, types.StartMessage)
	} else if *mode == "e" {
		sendMessage(ws, types.EndMessage)
	}
	
	/*if (*tail) {
		tailOutput(ws)
	}*/
}
