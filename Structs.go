package electrum

import (
	"encoding/json"
	"errors"
	"net"
	"strings"
	"time"
)

type ElectrumClient struct {
	Conn    net.Conn
	Network string
}

func NewElectrumClient(address, network string) (*ElectrumClient, error) {
	network = strings.ToLower(network)
	supported := isSupportedNetwork(network)
	if !supported {
		return nil, errors.New("unsupported network")
	}

	conn, err := net.DialTimeout("tcp", address, 7*time.Second)
	if err != nil {
		return nil, err
	}
	return &ElectrumClient{
		Conn:    conn,
		Network: networks[network],
	}, nil

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

// TxHistory represents a transaction history entry with transaction hash and height.
type TxHistory struct {
	TxHash string `json:"tx_hash"`
	Height int    `json:"height"`
}
