package controllers

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"net/http"
	"strconv"
	libeth "zmakers-backend/lib/eth"
	"zmakers-backend/models"
	"zmakers-backend/models/utility"
)

type NftController struct {
	BC    *BaseController
}


var privateKey = "6c7117111a42dd5dfcff752ee0b32c3f85699192d6c18297fc23d473bf8089c9"
var nftContract = "0xdE87CaE5166fBDbAd18581af77ab281DB115fD21"

var pics = []string{"http://rguehz22i.bkt.clouddn.com/zhaji1.png", "http://rguehz22i.bkt.clouddn.com/zhaji2.png", "http://rguehz22i.bkt.clouddn.com/zhaji3.png", "http://rguehz22i.bkt.clouddn.com/zhaji4.png", "http://rguehz22i.bkt.clouddn.com/zhaji5.png", "http://rguehz22i.bkt.clouddn.com/kele.png", "http://rguehz22i.bkt.clouddn.com/nvpu.png"}

func (nc *NftController)  MintNft(w http.ResponseWriter, r *http.Request)  {
	var jsonResult utility.Response
	sk, nftOwner := libeth.GenKeyAddressPair()
	log.Info("nftOwner privateKey: %s", sk)
	rewardIndex := r.URL.Query().Get("picIndex")
	if rewardIndex == "" {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "请求参数不正确"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return

	}
	mintFuncSig := crypto.Keccak256([]byte("mint(address,uint256)"))[:4]
	param1 := common.LeftPadBytes(common.HexToAddress(nftOwner).Bytes(), 32)
	param2 := common.LeftPadBytes(common.FromHex("0x"+rewardIndex), 32)
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
	index, _ := strconv.Atoi(rewardIndex)
	tblNft := models.Tbl_Nft{
		TxHash:   txHash,
		Address:  nftOwner,
		TokenURI: pics[index],
	}
	err = tblNft.CreatNft()
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "create Nft failed"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Data = map[string]string{"txHash": txHash}
	jsonResult.Code = http.StatusOK
	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}

func (nc *NftController) GetTokenInfo(w http.ResponseWriter, r *http.Request) {
	txHash := r.URL.Query().Get("txHash")
	var jsonResult utility.Response
	nft := &models.Tbl_Nft{
		TxHash:   txHash,
	}
	nft = nft.GetNftByTxHash()
	if nft == nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "区块链网络拥堵，nft生成失败，请稍后再试"}
		nc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Data = nft
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
	//fmt.Printf("tokenUri %s\n", string(res[32:]))
	returnData := &ResultData {
		from,
		to,
		tokenId,
		string(res)[64:107],
	}

	jsonResult.Data = returnData
	nc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}
