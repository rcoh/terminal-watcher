package types

import (
	"encoding/json"
)
const Start_id = 0
const End_id = 1

type Message struct {
	Mode int
	Command string
	Status int
}

func StartMessage(command string) Message {
	return Message{
		Mode: Start_id,
		Command: command,
	}
}

func EndMessage(command string, status int) Message { 
	return Message{
		Mode: End_id,
		Command: command,
		Status: status,
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



