// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"github.com/rcoh/terminal-watcher/types"
)

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan connectionWithMessage

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	// Keep track of clients for ids
	idConnectionMap map[string]*connection
}

type connectionWithMessage struct {
	connection *connection
	message []byte
}

var h = hub{
	broadcast:   make(chan (connectionWithMessage)),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
	idConnectionMap: make(map[string]*connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:

			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case conMessage := <-h.broadcast:

			payload := conMessage.message
			incomingConnection := conMessage.connection
			message, err := types.Deserialize(payload)

			if err != nil {
				log.Println("Error deserializing: ", err)
				continue
			}
			log.Println(message)

			switch message.Mode {
			case types.Listen_id:
				if (message.ClientId != "") {
					log.Println("Registering listener: ", message.ClientId)
					h.idConnectionMap[message.ClientId] = incomingConnection
				} else {
					log.Println("Got bad message (client id not set)")
				}

			default:
				// Send the message to the clients listening on that id
				// TODO: enforce every message has a client id
				id := message.ClientId
				outgoingConnection, ok := h.idConnectionMap[id]
				_, validChannel := h.connections[outgoingConnection]
				if ok && validChannel {
					select {
						case outgoingConnection.send <- payload:
							log.Println("message routed")
						default:
							println("listener dead, closing")
							close(outgoingConnection.send)
							delete(h.idConnectionMap, id)

					}
				} else {
					log.Println("Got", message.ClientId, "but no client registered")
				}
			}
		}
	}
}