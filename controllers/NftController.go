package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	nftOwner := r.URL.Query().Get("nftOwner")
	fmt.Printf("nftOwner %s\n", nftOwner)
	mintFuncSig := crypto.Keccak256([]byte("mint(address)"))[:4]
	_nftOwner := common.LeftPadBytes(common.HexToAddress(nftOwner).Bytes(), 32)
	data := append(mintFuncSig, _nftOwner...)
	fmt.Printf("data %s", data)
	tx := libeth.TransactionObj{
		"0xf09f2d5c6c57088173A0280a7F40ABA486585B92",
		int64(0),
		[]byte(data),
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
	jsonResult.Data = map[string]string{"txHash": txHash}
	jsonResult.Code = http.StatusOK

	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}

func (nc *NftController)  ParseTx(w http.ResponseWriter, r *http.Request) {
	txHash := r.URL.Query().Get("txHash")
	fmt.Printf("txHash %s\n", txHash)
	receipt, err := libeth.GetTransactionReceipt(txHash)
	fmt.Printf("receipt %v\n", receipt.Status)
	var jsonResult utility.Response
	if receipt == nil || err != nil || receipt.Status != 1 {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": err.Error()}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	log := receipt.Logs[0]
	from := "0x"+log.Topics[1].String()[26:]
	to := "0x"+log.Topics[2].String()[26:]
	tokenId := libeth.Hex2Dec(log.Topics[3].String()[2:])
	//tokenURISig := crypto.Keccak256([]byte("tokenURI"))[:4]
	//msg := append(tokenURISig, []byte(string(tokenId))...)
	//libeth.CallContract(msg)

	jsonResult.Code = http.StatusOK
	type ResultData struct {
		From string `json:"from"`
		To string `json:"to"`
		TokenId int `json:"tokenId"`
	}
	returnData := &ResultData {
		from,
		to,
		tokenId,
	}
	resultJson, _ := json.Marshal(returnData)
	jsonResult.Data = string(resultJson)
	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}
