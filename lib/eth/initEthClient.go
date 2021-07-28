package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

var (
	ethClient *ethclient.Client
	chainId *big.Int
)

// init ethClient
func init() {
	ethClient = initEthClient()
	getChainId()
}

func initEthClient() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		os.Exit(1)
	}
	return client
}


func getChainId() {
	var err error
	chainId, err = ethClient.ChainID(context.Background())
	if err != nil {
		os.Exit(2)
	}
	fmt.Println("chainId: ", chainId)
}

