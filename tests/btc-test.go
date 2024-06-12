package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/crytoken/electrum-go"
)

func main() {

	const electrumServer = "bolt.schulzemic.net:50001"
	address := "bc1q9x4jmpf3ptxm52fpdm46cg8r24r088tzqqmwpg"
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

	_, status3, err := e.Subscribe("bc1q9x4jmpf3ptxm52fpdm46cg8r24r088tzqqmwpg")
	if err != nil {
		fmt.Println("errc", err)
		os.Exit(1)
	}
	fmt.Println("Status 3", status3)

	txHistory, addr, err := e.GetTxHistory(address)
	if err != nil {
		fmt.Println("Err history:", err)
		os.Exit(1)
	}
	fmt.Println("History for: ", addr)
	historyLen := len(txHistory)
	for indx, tx := range txHistory {
		fmt.Printf("%d.TxHash: %s,TxHeight:%v\n", indx+1, tx.TxHash, tx.Height)
		if indx == historyLen-1 {
			trans, err := e.GetTx(tx.TxHash, true)
			if err != nil {
				fmt.Printf("GetTx Err: %v", err)
			}
			fmt.Println(trans.Confirmations)
		}
	}

	tx, err := e.GetTx("86748672613b248dc622ab64a1e4b67c112fcf6072f056602a7852e576be4676", true)
	if err != nil {
		fmt.Printf("GetTx Err: %v", err)
	}
	fmt.Printf("Transaction: %+v\n", tx)
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
