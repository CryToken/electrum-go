package electrum

import (
	"crypto/sha256"
	"encoding/hex"
)

func getScriptHash(scriptPubKey string) string {
	scriptBytes, _ := hex.DecodeString(scriptPubKey)
	hash := sha256.Sum256(scriptBytes)
	hashHex := hex.EncodeToString(hash[:])
	// Convert to Electrum's endianness
	var result string
	for i := len(hashHex) - 2; i >= 0; i -= 2 {
		result += hashHex[i : i+2]
	}
	return result
}
