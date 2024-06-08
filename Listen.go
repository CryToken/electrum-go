package electrum

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

//You can run it in goroutin to Listen Notifications and do not block main Threat.

func (s *ElectrumServer) ListenForNotification() {
	ticker := time.NewTicker(3 * time.Minute)

	defer ticker.Stop()
	reader := bufio.NewReader(s.Conn)

	for {
		select {
		case <-ticker.C:
			err := s.SendPing()
			if err != nil {
				log.Printf("Failed to send ping: %v", err)
			} else {
				log.Println("Ping sent successfully")
			}

		default:
			responseData, err := reader.ReadBytes('\n')
			if err != nil {
				log.Printf("Error reading response: %v", err)
				continue
			}

			var notification struct {
				Method string        `json:"method"`
				Params []interface{} `json:"params"`
			}

			err = json.Unmarshal(responseData, &notification)
			if err != nil {
				log.Printf("Failed to unmarshaling Notify: %v", err)
				continue
			}
			if notification.Method == "blockchain.scripthash.subscribe" {
				if len(notification.Params) == 2 {
					scriptHash := notification.Params[0].(string)
					status := notification.Params[1].(string)
					//Change logic for your own purpose.
					fmt.Printf("Received notification.\n[Script Hash: %s,New Status: %s]\n\n", scriptHash, status)
				} else {
					fmt.Println("Unexpected notification format:", string(responseData))
				}
			}

		}
	}
}
