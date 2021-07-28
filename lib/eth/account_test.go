package eth

import (
	"fmt"
	"testing"
)

func TestGenAcc(t *testing.T) {
	sk, address := genKeyAddressPair()
	fmt.Println(sk)
	fmt.Println(address)
}
