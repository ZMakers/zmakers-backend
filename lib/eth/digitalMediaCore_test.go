package eth

import (
	"fmt"
	"testing"
)

func TestUnPause(t *testing.T) {
	txHash, err := tx.unPause()
	fmt.Printf("txHash %v, err: %v\n", txHash, err)
}

func TestPause(t *testing.T) {
	txHash, err := tx.pause()
	fmt.Printf("txHash %v, err: %v\n", txHash, err)
}

func TestSetV1dmsAddress(t *testing.T) {
	v1DmsAddress := "0xBc37d5e0CCE1Ebb8cDE9DF3Ec2FB0b8a9628f8F0"
	txHash, err := tx.setV1DigitalMediaStoreAddress(v1DmsAddress)
	fmt.Printf("txHash %v, err: %v\n", txHash, err)
}

func TestCreateCollection(t *testing.T) {
	metaPath := "https://dpos.com"
	txHash, err := tx.createCollection(metaPath)
	fmt.Printf("txHash %v, err: %v\n", txHash, err)
}

var callDMC = callObj{
	To:      "",
	Value:   0,
	Data:    nil,
	PrivKey: "",
}

func TestGetCollection(t *testing.T) {
	res, err := callDMC.getCollection("4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestAddApprovedTokenCreator(t *testing.T) {
	creator := "0xb9f745C18183c7a46bF80F00Aa402d229a7d201d"
	txHash, err := tx.addApprovedTokenCreator(creator)
	fmt.Printf("txHash %v, err: %v\n", txHash, err)
}

