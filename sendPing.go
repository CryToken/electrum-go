package electrum

import (
	"encoding/json"
	"fmt"
)

func (s *ElectrumClient) SendPing() error {
	const method = "server.ping"
	request := ElectrumRequest{
		ID:      1,
		Method:  method,
		Params:  []interface{}{},
		Jsonrpc: "2.0",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling ping request: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return fmt.Errorf("error writing ping request: %w", err)
	}

	return nil
}
