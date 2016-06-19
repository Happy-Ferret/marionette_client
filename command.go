package marionette_client

import (
	"encoding/json"
)

func makeProto2Command(command string, values interface{}) ([]byte, error) {
	message := make(map[string]interface{})
	message["name"] = command
	message["parameters"] = values
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func makeProto3Command(msgID int, command string, values interface{}) ([]byte, error) {
	message := make([]interface{}, 4)
	message[0] = 0
	message[1] = msgID
	message[2] = command
	message[3] = values
	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(b)) //debug only.

	return b, nil
}
