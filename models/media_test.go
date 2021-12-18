package models

import (
	"fmt"
	"testing"
	"zmakers-backend/service"
)

func TestCreateMedia(t *testing.T) {
	//id, _ := uint256.FromHex("0x1")
	//collectionId, _ := uint256.FromHex("0x1")
	m := Media{
	14,
	5,
	1,
	"https://asdasd.com",
	3,
	"0xasqade",
	}
	service.DbInit()
	fmt.Println(m.CreateMedia())

}


func TestBuyMedia(t *testing.T) {
	transfer := Transfer{
		TxHash:  "0x123weq",
		From:    "0x14443",
		To:      "0x13121",
		MediaId: 0,
	}
	service.DbInit()

	transfer.BuyMedia()
}
