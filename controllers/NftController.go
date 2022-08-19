package controllers

import (
	"encoding/json"
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
var nftContract = "0xdE87CaE5166fBDbAd18581af77ab281DB115fD21"
func (nc *NftController)  MintNft(w http.ResponseWriter, r *http.Request)  {
	var jsonResult utility.Response

	nftOwner := r.URL.Query().Get("nftOwner")
	rewardIndex := r.URL.Query().Get("picIndex")
	if nftOwner == "" || rewardIndex == "" {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "请求参数不正确"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return

	}
	mintFuncSig := crypto.Keccak256([]byte("mint(address,uint256)"))[:4]
	param1 := common.LeftPadBytes(common.HexToAddress(nftOwner).Bytes(), 32)
	param2 := common.LeftPadBytes([]byte(rewardIndex), 32)
	data := append(mintFuncSig, param1...)
	data = append(data, param2...)
	tx := libeth.TransactionObj{
		nftContract,
		int64(0),
		data,
		uint64(5100000),
		big.NewInt(15000),
		privateKey,
	}
	txHash, err := tx.SendTx()
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

func (nc *NftController) ParseMintNftTx(w http.ResponseWriter, r *http.Request) {
	var jsonResult utility.Response
	txHash := r.URL.Query().Get("txHash")
	if  txHash == "" {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "请求参数不正确"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return

	}

	receipt, err := libeth.GetTransactionReceipt(txHash)
	if receipt == nil || err != nil || receipt.Status != 1 {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "还未生成nft或者没有成功生成nft"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	log := receipt.Logs[0]
	from := "0x"+log.Topics[1].String()[26:]
	to := "0x"+log.Topics[2].String()[26:]
	tokenId := libeth.Hex2Dec(log.Topics[3].String()[2:])
	tokenURISig := crypto.Keccak256([]byte("tokenURI(uint256)"))[:4]
	param := common.LeftPadBytes([]byte(string(tokenId)),32)
	msg := append(tokenURISig, param...)
	res, err := libeth.CallContract(nftContract, msg)

	if err !=nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": err.Error()}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	type ResultData struct {
		From string `json:"from"`
		To string `json:"to"`
		TokenId int `json:"tokenId"`
		TokenUri string `json:"tokenUri"`
	}

	returnData := &ResultData {
		from,
		to,
		tokenId,
		string(res),
	}

	resultJson, _ := json.Marshal(returnData)
	jsonResult.Data = string(resultJson)
	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}
