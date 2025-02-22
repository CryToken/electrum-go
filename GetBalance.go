package electrum

import (
	"encoding/json"
	"fmt"
)

func (s *ElectrumClient) GetBalance(address string) (int64, int64, error) {
	const method = "blockchain.scripthash.get_balance"

	//Getting ScriptHash
	scriptHash, err := AddressToScriptHash(address)
	if err != nil {
		return 0, 0, err
	}
	var balance Balance

	request := ElectrumRequest{
		ID:      1,
		Method:  method,
		Params:  []interface{}{scriptHash},
		Jsonrpc: "2.0",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return 0, 0, fmt.Errorf("error marshaling request: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return 0, 0, fmt.Errorf("error writing request: %w", err)
	}

	buffer := make([]byte, 4096)
	n, err := s.Conn.Read(buffer)
	if err != nil {
		return 0, 0, fmt.Errorf("error reading response: %w", err)
	}

	var response ElectrumResponse
	err = json.Unmarshal(buffer[:n], &response)
	if err != nil {
		return 0, 0, fmt.Errorf("error unmarshaling response: %w", err)
	}

	if response.Error != nil {
		return 0, 0, fmt.Errorf("error from server: %s", response.Error.Message)
	}

	err = json.Unmarshal(response.Result, &balance)
	if err != nil {
		return 0, 0, fmt.Errorf("error unmarshaling balance: %w", err)
	}

	return balance.Confirmed, balance.Unconfirmed, nil
}
