package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func getBlockByNumber(number *big.Int)  (*types.Block, error) {
	block, err := ethClient.BlockByNumber(context.Background(), number)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return block, nil

}

