package eth

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func genKeyAddressPair() (sk string, address string) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Error("generate private key error: %v", err)
		return "", ""
	}
	skBytes := crypto.FromECDSA(key)
	sk = "0x"+hex.EncodeToString(skBytes)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	address = addr.String()
	return
}
