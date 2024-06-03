package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func (s *ElectrumServer) ListenForNotification() {
	reader := bufio.NewReader(s.Conn)
	for {
		responseData, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		var notification struct {
			Method string        `json:"method"`
			Params []interface{} `json:"params"`
		}

		err = json.Unmarshal(responseData, &notification)
		if err != nil {
			fmt.Println("Error unmarshalling notification:", err)
			continue
		}

		if notification.Method == "blockchain.scripthash.subscribe" {
			if len(notification.Params) == 2 {
				scriptHash := notification.Params[0].(string)
				status := notification.Params[1].(string)
				fmt.Printf("Received notification.\n[Script Hash: %s,New Status: %s]\n\n", scriptHash, status)
			} else {
				fmt.Println("Unexpected notification format:", string(responseData))
			}
		}
	}
}
