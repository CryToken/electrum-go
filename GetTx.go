package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func (s *ElectrumServer) GetTx(txid string, verbose bool) (*Tx, error) {
	const method = "blockchain.transaction.get"

	request := ElectrumRequest{
		ID:      5,
		Method:  method,
		Params:  []interface{}{txid, verbose},
		Jsonrpc: "2.0",
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	_, err = s.Conn.Write(append(requestData, '\n'))
	if err != nil {
		return nil, fmt.Errorf("error writing request: %w", err)
	}

	reader := bufio.NewReader(s.Conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Попытка разобрать ответ как JSON
	var jsonResponse map[string]interface{}
	err = json.Unmarshal([]byte(response), &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Преобразование ответа в структуру Tx
	result, ok := jsonResponse["result"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	txBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error marshaling result to JSON: %w", err)
	}

	var tx Tx
	err = json.Unmarshal(txBytes, &tx)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling result to Tx struct: %w", err)
	}

	return &tx, nil
}

type ScriptPubKey struct {
	Addresses []string `json:"addresses"`
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs"`
	Type      string   `json:"type"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vin struct {
	Txid        string    `json:"txid"`
	Vout        int       `json:"vout"`
	ScriptSig   ScriptSig `json:"scriptSig"`
	Sequence    uint32    `json:"sequence"`
	Txinwitness []string  `json:"txinwitness"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type Tx struct {
	BlockHash     string `json:"blockhash,omitempty"`
	BlockTime     int    `json:"blocktime,omitempty"`
	Confirmations int    `json:"confirmations,omitempty"`
	Hash          string `json:"hash"`
	Hex           string `json:"hex,omitempty"`
	Locktime      int    `json:"locktime"`
	Size          int    `json:"size"`
	Time          int    `json:"time,omitempty"`
	Txid          string `json:"txid"`
	Version       int    `json:"version"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	Vsize         int    `json:"vsize,omitempty"`
	Weight        int    `json:"weight,omitempty"`
}
