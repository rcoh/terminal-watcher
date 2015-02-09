package types

import (
	"encoding/json"
)

const Start_id = 0
const End_id = 1
const Listen_id = 2
const Tail_id = 3

type Message struct {
	ClientId string
	Mode int
	Command string
	Status int
}

// TODO: builder to avoid having to pass clientId

func StartMessage(command string, clientId string) Message {
	return Message{
		Mode: Start_id,
		Command: command,
		ClientId: clientId,
	}
}

func EndMessage(command string, status int, clientId string) Message { 
	return Message{
		Mode: End_id,
		Command: command,
		Status: status,
		ClientId: clientId,
	}
}

func TailMessage(command string, clientId string) Message {
	return Message{
		Mode: Tail_id,
		Command: command,
		ClientId: clientId,
	}
}

func Deserialize(input []byte) (*Message, error) {
	var message = &Message{}
	err := json.Unmarshal(input, message)
	return message, err
}

func (message *Message) Serialize() ([]byte, error) {
	return json.Marshal(message);
}



