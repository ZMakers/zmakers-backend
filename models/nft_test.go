package models

import (
	"log"
	"testing"
	"zmakers-backend/service"
)

func TestCreateNft(t *testing.T)  {
	service.DbInit()
	testNfts := [...]Tbl_Nft{
		{
			4,
			"0x0b7d65b8023e360bb84569a0e80e349269a16300dacc543a8f9f9a574fe6748c",
			"0x78C531c538A5013121E81BBA80cB9D3Bcee46962",
			"http://rguehz22i.bkt.clouddn.com/zhaji5.png",
		},
	}
	for _, testnft := range testNfts{
		err := testnft.CreatNft()
		if err != nil {
			log.Printf("err: %s", err)
		}
	}
}