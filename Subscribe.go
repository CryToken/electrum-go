package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
)

// Then you Subscribe to scriptHash,server return Current STATUS.Then something change you will get new status and scriptHash.
func (s *ElectrumClient) Subscribe(address string) (string, string, error) {

	//Getting ScriptHash
	scriptHash, err := AddressToScriptHash(address)
	if err != nil {
		return "", "", err
	}

	//Request param
	request := ElectrumRequest{
		ID:     3,
		Method: "blockchain.scripthash.subscribe",
		Params: []interface{}{scriptHash}, Jsonrpc: "2.0",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}
	reader := bufio.NewReader(s.Conn)

	responseData, err := reader.ReadBytes('\n')
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}

	var response ElectrumResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}

	if response.Error != nil {
		return scriptHash, "", fmt.Errorf("error: %s", response.Error.Message)
	}

	var status string
	err = json.Unmarshal(response.Result, &status)
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}

	return scriptHash, status, nil
}
