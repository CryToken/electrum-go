package electrum

import (
	"strings"
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

		return true

	}
	if strings.HasPrefix(address, "1") || strings.HasPrefix(address, "3") {

		return true

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
