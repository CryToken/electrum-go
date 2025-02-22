package electrum

import (
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/bech32"
)

var networks = map[string]string{"bitcoin": "Bitcoin", "litecoin": "Litecoin"}

func isSupportedNetwork(network string) bool {
	_, exist := networks[network]
	if exist {
		return exist
	}
	return false
}

func isBitcointAddress(address string) bool {
	if strings.HasPrefix(address, "bc1") {
		_, _, err := bech32.Decode(address)
		if err == nil {
			return true
		}

	}
	if strings.HasPrefix(address, "1") || strings.HasPrefix(address, "3") {
		_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
		if err == nil {
			return true
		}

	}
	return false

}

func isLitecoinAddrees(address string) bool {
	if strings.HasPrefix(address, "ltc1") {
		return true
	}

	if strings.HasPrefix(address, "L") {
		return true
	}
	if strings.HasPrefix(address, "3") {
		return true
	}
	return false
}

func reverseHash(hash *[32]byte) {
	for i, j := 0, len(hash)-1; i < j; i, j = i+1, j-1 {
		hash[i], hash[j] = hash[j], hash[i]
	}
}
