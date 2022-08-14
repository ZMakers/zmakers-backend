package controllers

import (
	"math/big"
	"net/http"
	libeth "zmakers-backend/lib/eth"
	"zmakers-backend/models/utility"
)

type NftController struct {
	BC    *BaseController
}
var privateKey = "6c7117111a42dd5dfcff752ee0b32c3f85699192d6c18297fc23d473bf8089c9"
func (nc *NftController)  MintNft(w http.ResponseWriter, r *http.Request)  {
	nftOwner := r.Header.Get("nftOwner")
	tx := libeth.TransactionObj{
		"0xeD774722eC7fBEE84117D7da7856a1bC8a60B813",
		int64(0),
		[]byte(nftOwner),
		uint64(5100000),
		big.NewInt(15000),
		privateKey,
	}
	txHash, err := tx.SendTx()
	var jsonResult utility.Response
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": err.Error()}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	jsonResult.Data = txHash
	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}
