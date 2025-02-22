package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func (s *ElectrumClient) GetTxHistory(address string) ([]TxHistory, string, error) {
	const method = "blockchain.scripthash.get_history"

	//Getting ScriptHash
	scriptHash, err := AddressToScriptHash(address)
	if err != nil {
		return []TxHistory{}, "", err
	}

	var history []TxHistory

	request := ElectrumRequest{
		ID:      2,
		Method:  method,
		Params:  []interface{}{scriptHash},
		Jsonrpc: "2.0",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, "", fmt.Errorf("error marshaling request: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return nil, "", fmt.Errorf("error writing request: %w", err)
	}

	response, err := bufio.NewReader(s.Conn).ReadString('\n')
	if err != nil {
		return nil, "", fmt.Errorf("error reading response: %w", err)
	}

	var electrumResponse ElectrumResponse
	err = json.Unmarshal([]byte(response), &electrumResponse)
	if err != nil {
		return nil, "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	if electrumResponse.Error != nil {
		return nil, "", fmt.Errorf("error from server: %s", electrumResponse.Error.Message)
	}

	err = json.Unmarshal(electrumResponse.Result, &history)
	if err != nil {
		return nil, "", fmt.Errorf("error unmarshaling history: %w", err)
	}

	return history, address, nil
}
