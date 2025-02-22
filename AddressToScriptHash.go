package electrum

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"

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
