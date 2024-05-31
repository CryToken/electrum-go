package electrum

import (
	"encoding/json"
	"fmt"
	"net"
)

func Ping(conn net.Conn) error {

	request := ElectrumRequest{
		ID:      1,
		Method:  "server.ping",
		Params:  []interface{}{},
		Jsonrpc: "2.0",
	}

	//Marshal Request to JSON
	requestData, err := json.Marshal(request)
	if err != nil {
		fmt.Errorf("error marshaling request: %w", err)
	}

	// Send the request
	_, err = conn.Write(append(requestData, '\n'))
	if err != nil {
		fmt.Errorf("error Sending request: %w", err)
	}

	// Read the response
	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Errorf("error Readimg response: %w", err)
	}

	// Unmarshal the response
	var response ElectrumResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		fmt.Errorf("error Unmarshaling request: %w", err)
	}

	// Check for errors in the response
	if response.Error != nil {
		fmt.Errorf("error from Server: %w", response.Error.Message)
	}

	return nil
}
