package types

import (
	"encoding/json"
)
const Start_id = 0
const End_id = 1

type Message struct {
	Mode int
}

var StartMessage = Message{
	Mode: Start_id,
}

var EndMessage = Message{
	Mode: End_id,
}

func Deserialize(input []byte) (*Message, error) {
	var message = &Message{}
	err := json.Unmarshal(input, message)
	return message, err
}

func (message *Message) Serialize() ([]byte, error) {
	return json.Marshal(message);
}



