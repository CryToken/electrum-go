electrum-go 
===

It's a library to work with Electrum Server(local or remote).

## Install 

```bash
$ go get github.com/crytoken/electum-go@latest 
```

## Getting Started

``` go
package main

import (
	"fmt"
	"log"

	"github.com/crytoken/electrum-go"
)

func main() {

	client, err := electrum.NewElectrumClient("bitcoin.aranguren.org:50001", "bitcoin")
	if err != nil {
		panic(err)
	}

	conf, unconf, err := client.GetBalance("bc1pxc3entkl3v09ggcfe9nvcuq720plfu4lf5frm3yw0a39zckuasksl83a2s")
	if err != nil {
		panic(err)
	}

	fmt.Println(conf, unconf)

	tx, err := client.GetTx("625c8570597028fc810538d873371f6a086b7952efb96e458065a78d3ba64f1d", true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Tx : %+v\n", tx)

	
	script, status, err := client.Subscribe("bc1pxc3entkl3v09ggcfe9nvcuq720plfu4lf5frm3yw0a39zckuasksl83a2s")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Scrpt:", script, "Status:", status)
	client.ListenForNotification()

}
```

You can:
 GetBalance()
 
 GetTxHistory()
 
 Subscribe for Changes


Support only BItcoin and Litecoin Networks right now.

