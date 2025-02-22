package electrum

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil/bech32"

	ltcchaincfg "github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/ltcutil"
	ltctxscript "github.com/ltcsuite/ltcd/txscript"
)

func AddressToScriptHash(address string) (string, error) {
	if isBitcointAddress(address) {
		addr, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
		if err != nil {
			return "", err

		}
		script, err := txscript.PayToAddrScript(addr)
		if err != nil {
			return "", err
		}
		scritHash := sha256.Sum256(script)
		reverseHash(&scritHash)
		return hex.EncodeToString(scritHash[:]), nil
	}

	if isLitecoinAddrees(address) {
		addr, err := ltcutil.DecodeAddress(address, &ltcchaincfg.MainNetParams)
		if err != nil {
			return "", err
		}

		script, err := ltctxscript.PayToAddrScript(addr)
		if err != nil {
			return "", err
		}
		scriptHash := sha256.Sum256(script)
		reverseHash(&scriptHash)
		return hex.EncodeToString(scriptHash[:]), nil
	}

	return "", errors.New("unsupported address format")
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
