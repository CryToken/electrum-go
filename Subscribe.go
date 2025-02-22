package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
)

// Then you Subscribe to scriptHash,server return Current STATUS.Then something change you will get new status and scriptHash.
func (s *ElectrumServer) Subscribe(address string) (string, string, error) {

	var scriptPubKey string
	var err error

	if s.Network == "Bitcoin" {
		scriptPubKey, err = getScriptPubKey(address)
		if err != nil {
			return "", "", fmt.Errorf("error getting scriptPubKey: %w", err)
		}
	} else if s.Network == "Litecoin" {
		scriptPubKey, err = getLitecoinScriptPubKey(address)
		if err != nil {
			return "", "", fmt.Errorf("error getting scriptPubKey: %w", err)
		}
	} else {
		return "", "", fmt.Errorf("Unsupported Network: %s", s.Network)
	}

	//Getting ScriptHash
	scriptHash := getScriptHash(scriptPubKey)
	//Request param
	request := ElectrumRequest{
		ID:     3,
		Method: "blockchain.scripthash.subscribe",
		Params: []interface{}{scriptHash}, Jsonrpc: "2.0",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return scriptHash, "", fmt.Errorf("Error: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return scriptHash, "", fmt.Errorf("Error: %w", err)
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
		return scriptHash, "", fmt.Errorf("error: %s", response.Error)
	}

	var status string
	err = json.Unmarshal(response.Result, &status)
	if err != nil {
		return scriptHash, "", fmt.Errorf("error: %w", err)
	}

	return scriptHash, status, nil
}
