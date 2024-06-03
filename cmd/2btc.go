package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/crytoken/electrum-go"
)

func main() {

	const electrumServer = "exs.dyshek.org:50001"
	address := "3NLezRBqRYYGBvNdQSW9etkvAb5f3fr1Kp"
	conn, err := net.DialTimeout("tcp", electrumServer, 5*time.Second)
	if err != nil {
		fmt.Println("Failed to connect to Electrum server:", err)
		return
	}
	defer conn.Close()

	e := electrum.ElectrumServer{Conn: conn, Network: "Bitcoin"}
	conf, unconf, err := e.GetBalance(address)
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}

	fmt.Println(satoshiToBTC(conf), unconf)
	_, status, err := e.Subscribe(address)
	if err != nil {
		fmt.Println("errc", err)
		os.Exit(1)
	}
	fmt.Println(status)

	_, status2, err := e.Subscribe("321mDQoLvBc1qewWkyaG5286m81EivreUj")
	if err != nil {
		fmt.Println("errc", err)
		os.Exit(1)
	}
	fmt.Println("Status 2", status2)

	_, status3, err := e.Subscribe("3LfmsZxY89VzPtgNRSoyiwFn61ayyaCaZ9")
	if err != nil {
		fmt.Println("errc", err)
		os.Exit(1)
	}
	fmt.Println("Status 3", status3)

	txHistory, addr, err := e.GetTxHistory("bc1q8yj0herd4r4yxszw3nkfvt53433thk0f5qst4g")
	if err != nil {
		fmt.Println("Err history:", err)
		os.Exit(1)
	}
	fmt.Println("History for: ", addr)
	for indx, tx := range txHistory {
		fmt.Printf("%d.TxHash: %s,TxHeight:%v\n", indx+1, tx.TxHash, tx.Height)
	}

	fmt.Println("Start Waiting Sub...")
	e.ListenForNotification()

}

func satoshiToBTC(satoshis int64) string {
	const satoshisInBTC = 100_000_000

	// Целая часть биткоинов
	btc := satoshis / satoshisInBTC
	// Дробная часть (сатоши)
	satoshiPart := satoshis % satoshisInBTC

	// Форматирование результата
	result := fmt.Sprintf("%d,%08d", btc, satoshiPart)

	return result
}
