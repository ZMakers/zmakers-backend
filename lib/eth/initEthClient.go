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
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/a2816bc61057481188c0a772cee3134e")
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

