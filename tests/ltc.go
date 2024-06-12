package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/crytoken/electrum-go"
)

func main() {

	electrumServer := "electrum-ltc.bysh.me:50002"
	address := "LiQCSrp6ATYHLr54SyQ5sJGPWRYmxYvj4M"

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // For testing purposes, don't verify the server's certificate
	}
	conn, err := tls.Dial("tcp", electrumServer, tlsConfig)
	if err != nil {
		fmt.Println("Failed to connect to Electrum server:", err)
		time.Sleep(5 * time.Second)

	}

	defer conn.Close()
	s := electrum.ElectrumServer{Conn: conn, Network: "Litecoin"}

	conf, unconf, err := s.GetBalance(address)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Balance:", conf, unconf)

	scriptHash, status, err := s.Subscribe(address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("ScriptHash:%s,status:%s\n", scriptHash, status)
	tx, err := s.GetTx("47dcd60f0b6760e8f335249fd26ae4359ab4fefab0ffb6d671d695b3d5e7c910", true)
	if err != nil {
		fmt.Printf("GetTx Err: %v", err)
	}
	x := tx.Vout[0].Value
	fmt.Printf("%T:%v\n", x, x)
	fmt.Printf("Transaction: %+v\n", tx)
	go s.ListenForNotification()
	select {}

}
