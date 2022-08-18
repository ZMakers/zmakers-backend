package eth

import (
	"fmt"
	"testing"
)

func TestGenAcc(t *testing.T) {
	sk, address := GenKeyAddressPair()
	fmt.Println(sk)
	fmt.Println(address)
}
