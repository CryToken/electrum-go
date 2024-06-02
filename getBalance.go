package electrum

import (
	"encoding/json"
	"fmt"
	"net"
)

// GetBalance retrieves the confirmed and unconfirmed balance for a given address from an Electrum server.
func GetBalance(conn net.Conn, address string) (int64, int64, error) {
	const method = "blockchain.scripthash.get_balance"

	// Getting scriptPubKey from address to get ScriptHash. It's required for Electrum Server.
	scriptPubKey, err := getScriptPubKey(address)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting scriptPubKey: %w", err)
	}

	// Getting scriptHash before sending the request.
	scriptHash := getScriptHash(scriptPubKey)

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

	_, err = conn.Write(append(requestData, '\n'))
	if err != nil {
		return 0, 0, fmt.Errorf("error writing request: %w", err)
	}

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
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

func GetLitecoinBalance(conn net.Conn, address string) (int64, int64, error) {
	const method = "blockchain.scripthash.get_balance"

	// Getting scriptPubKey from address to get ScriptHash. It's required for Electrum Server.
	scriptPubKey, err := getLitecoinScriptPubKey(address)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting scriptPubKey: %w", err)
	}

	// Getting scriptHash before sending the request.
	scriptHash := getScriptHash(scriptPubKey)

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

	_, err = conn.Write(append(requestData, '\n'))
	if err != nil {
		return 0, 0, fmt.Errorf("error writing request: %w", err)
	}

	buffer := make([]byte, 4096)
	n, err := conn.Read(buffer)
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

// Balance represents the balance structure with confirmed and unconfirmed fields.
type Balance struct {
	Confirmed   int64 `json:"confirmed"`
	Unconfirmed int64 `json:"unconfirmed"`
}

// ElectrumRequest represents the structure of an Electrum request.
type ElectrumRequest struct {
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Jsonrpc string        `json:"jsonrpc"`
}

// ElectrumResponse represents the structure of an Electrum response.
type ElectrumResponse struct {
	ID     int             `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  *ElectrumError  `json:"error,omitempty"`
}

// ElectrumError represents the structure of an error returned by the Electrum server.
type ElectrumError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
