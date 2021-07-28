package eth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGetBlockHeight(t *testing.T) {
	block, err := getBlockByNumber(big.NewInt(1))
	assert.Nil(t, err)
	fmt.Println(block.Hash())
}

