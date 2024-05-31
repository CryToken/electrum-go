package electrum

import (
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/bech32"
)

func getScriptPubKey(address string) (string, error) {
	if strings.HasPrefix(address, "bc1") {
		_, data, err := bech32.Decode(address)
		if err != nil {
			return "", err
		}
		converted, err := bech32.ConvertBits(data[1:], 5, 8, false)
		if err != nil {
			return "", err
		}
		var scriptPubKey []byte
		if data[0] == 0x00 { // P2WPKH
			scriptPubKey = append([]byte{0x00, byte(len(converted))}, converted...)
		} else { // Other SegWit versions
			scriptPubKey = append([]byte{data[0], byte(len(converted))}, converted...)
		}
		result := fmt.Sprintf("%x", scriptPubKey)
		return result, nil
	} else if strings.HasPrefix(address, "3") {
		decodedAddress := DecodeBase58(address)
		if len(decodedAddress) != 25 {
			return "", fmt.Errorf("invalid address length")
		}
		scriptHash := decodedAddress[1:21]
		scriptPubKey := append([]byte{0xa9, 0x14}, scriptHash...)
		scriptPubKey = append(scriptPubKey, 0x87)
		result := fmt.Sprintf("%x", scriptPubKey)

		return result, nil
	} else if strings.HasPrefix(address, "1") {
		decodedAddress := DecodeBase58(address)
		if len(decodedAddress) != 25 {
			return "", fmt.Errorf("invalid address length")
		}
		pubKeyHash := decodedAddress[1:21]
		scriptPubKey := append([]byte{0x76, 0xa9, 0x14}, pubKeyHash...)
		scriptPubKey = append(scriptPubKey, 0x88, 0xac)
		result := fmt.Sprintf("%x", scriptPubKey)
		return result, nil
	} else {
		return "", fmt.Errorf("unsupported address type")
	}

}

func DecodeBase58(b string) []byte {
	alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	result := []byte{}
	for _, c := range b {
		charIndex := -1
		for i, a := range alphabet {
			if a == c {
				charIndex = i
				break
			}
		}
		if charIndex == -1 {
			return nil
		}
		carry := charIndex
		for j := 0; j < len(result); j++ {
			carry += 58 * int(result[j])
			result[j] = byte(carry % 256)
			carry /= 256
		}
		for carry > 0 {
			result = append(result, byte(carry%256))
			carry /= 256
		}
	}
	// Reverse the byte slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	// Append leading zero bytes for each leading '1' character
	for i := 0; i < len(b) && b[i] == '1'; i++ {
		result = append([]byte{0}, result...)
	}
	return result
}
