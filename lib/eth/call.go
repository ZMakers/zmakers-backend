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

	return ethClient.CallContract(context.Background(), callMsg, nil)
}

func CallContract(contractAdd string, data []byte)  ([]byte, error)  {
	toAdd := common.HexToAddress(contractAdd)

	callMsg := ethereum.CallMsg{
		To: &toAdd,
		Data: data,
	}
	return ethClient.CallContract(context.Background(), callMsg, nil)

}