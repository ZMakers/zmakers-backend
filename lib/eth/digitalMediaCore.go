package eth

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strconv"
)

const dmcContract = "0x344b79B7cA266B8608a60c9AaC1139b88d50784c"
const dmsContract = "0xEB77CA33f7219edE12fbC8c8f282f06634F27B2e"
const dmsV1Contract =  "0x3ca7202013Dc434FDd7686576ce642fB2CE539e4"
const crsContract = "0x1D42B5f0a0a19E82D354706BE9Df863F9fcb2a33"

func (dmcTx *TransactionObj) invokeDmcContract(data []byte) (string, error) {
	dmcTx.To = dmcContract
	dmcTx.Value = 0
	dmcTx.Data = data

	return dmcTx.sendTx()
}

func (dmcTx *TransactionObj) unPause() (string, error) {
	return dmcTx.invokeDmcContract(crypto.Keccak256([]byte("unpause()"))[:4])
}

func (dmcTx *TransactionObj) pause() (string, error) {
	return dmcTx.invokeDmcContract(crypto.Keccak256([]byte("pause()"))[:4])
}

func (dmcTx *TransactionObj) setV1DigitalMediaStoreAddress(dmsAddress string) (string, error) {
	funcSignatureBytes := crypto.Keccak256([]byte("setV1DigitalMediaStoreAddress(address)"))[:4]
	_dmsAddress := common.LeftPadBytes(common.HexToAddress(dmsAddress).Bytes(), 32)
	return dmcTx.invokeDmcContract(append(funcSignatureBytes, _dmsAddress...))
}

func (dmcTx *TransactionObj) createCollection(metaPath string)  (string, error) {
	funcSignatureBytes := crypto.Keccak256([]byte("createCollection(string)"))[:4]
	metaPathHex := hex.EncodeToString([]byte(metaPath))
	fmt.Println(len(metaPathHex))
	fmt.Println(metaPathHex)
	metaPathBytes := common.LeftPadBytes(common.FromHex(metaPathHex), 32)
	return dmcTx.invokeDmcContract(append(funcSignatureBytes, metaPathBytes...))
}

type collection struct {
	Id int `json:"id"`
	Address string `json:"address"`
	MetaPath string `json:"metaPath"`
}

func (dmcCall *callObj) getCollection(collectionId string) (string, error) {
	dmcCall.To = dmcContract
	funcSignatureBytes := crypto.Keccak256([]byte("getCollection(uint256)"))[:4]

	_collectionId := new(big.Int)
	_collectionId.SetString(collectionId, 10)
	cId := common.LeftPadBytes(_collectionId.Bytes(), 32)

	dmcCall.Data = append(funcSignatureBytes, cId...)
	toAdd := common.HexToAddress(dmcCall.To)
	callMsg := ethereum.CallMsg{
		To: &toAdd,
		Data: dmcCall.Data,
	}

	_collection, err := ethClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return "", err
	}

	fmt.Println(len(_collection))

	cid := common.Bytes2Hex(_collection[:32])

	address := common.Bytes2Hex(_collection[44:64])
	fmt.Println(string(_collection[64:128]))

	Collection := collection{
		Id: Hex2Dec(cid),
		Address: "0x" + address,
		MetaPath: common.Bytes2Hex(_collection[64:]),
	}
	fmt.Println(Collection)

	return common.Bytes2Hex(_collection), nil
}

func (dmcTx *TransactionObj) addApprovedTokenCreator(creator string) (string, error) {
	funcSignatureBytes := crypto.Keccak256([]byte("addApprovedTokenCreator(address)"))[:4]
	_creatorAddress := common.LeftPadBytes(common.HexToAddress(creator).Bytes(), 32)
	return dmcTx.invokeDmcContract(append(funcSignatureBytes, _creatorAddress...))
}

func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}

