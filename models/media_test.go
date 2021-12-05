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
	1,
	5,
	1,
	"https://asdasd.com",
	3,
	"0xasqade",
	}
	service.DbInit()
	fmt.Println(m.CreateMedia())

}
