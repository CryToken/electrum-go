package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/crytoken/electrum-go"
)

func main() {
	const electrumServer = "exs.dyshek.org:50001"
	address := "1CTHxp8aoKRXNsEp1CBpuYq5Fe9daYEBDy"
	conn, err := net.DialTimeout("tcp", electrumServer, 5*time.Second)
	if err != nil {
		fmt.Println("Failed to connect to Electrum server:", err)
		return
	}
	defer conn.Close()

	err = electrum.Ping(conn)
	if err != nil {
		log.Fatal(err)
	}
	conf, unconf, err := electrum.GetBalance(conn, address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conf, unconf)

	txHistory, addr, _ := electrum.GetTxHistory(conn, address)
	fmt.Println("History for address: ", addr)
	for indx, Tx := range txHistory {
		fmt.Print(indx + 1)
		fmt.Printf("TxHash: %v\nHeight: %v\n", Tx.TxHash, Tx.Height)
	}
}
