package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type callObj struct {
	To string `json:"to"`
	Value int64 `json:"value"`
	Data []byte `json:"data"`
	PrivKey string `json:"privKey"`
}

func (coj *callObj)  callContract() ([]byte, error) {

	toAdd := common.HexToAddress(coj.To)
	callMsg := ethereum.CallMsg{
		To: &toAdd,
		Data: coj.Data,
	}

	result, err := ethClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}