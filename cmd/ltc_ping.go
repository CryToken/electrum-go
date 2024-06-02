package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/crytoken/electrum-go"
)

func main() {
	electrumServer := "electrum-ltc.bysh.me:50002"
	address := "LYhttvnKawAv6RcHQ4eBkNtifuiEA99PFe"

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

	err = electrum.Ping(conn)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Server Online")

	conf, unconf, err := electrum.GetLitecoinBalance(conn, address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf, unconf)

}
