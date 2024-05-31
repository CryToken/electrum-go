package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

// GetTxHistory retrieves the transaction history for a given address from an Electrum server.
func GetTxHistory(conn net.Conn, address string) ([]TxHistory, string, error) {
	const method = "blockchain.scripthash.get_history"

	// Getting scriptPubKey from address to get ScriptHash. It's required for Electrum Server.
	scriptPubKey, err := getScriptPubKey(address)
	if err != nil {
		return nil, "", fmt.Errorf("error getting scriptPubKey: %w", err)
	}

	// Getting scriptHash before sending the request.
	scriptHash := getScriptHash(scriptPubKey)

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

	_, err = conn.Write(append(requestData, '\n'))
	if err != nil {
		return nil, "", fmt.Errorf("error writing request: %w", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
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

// TxHistory represents a transaction history entry with transaction hash and height.
type TxHistory struct {
	TxHash string `json:"tx_hash"`
	Height int    `json:"height"`
}
