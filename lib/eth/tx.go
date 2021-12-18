package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

type TransactionObj struct {
	To string `json:"to"`
	Value int64 `json:"value"`
	Data []byte `json:"data"`
	Gas uint64 `json:"gas"`
	GasPrice *big.Int `json:"gasPrice"`
	PrivKey string `json:"privKey"`
}

func calculatePendingNonce(sk *ecdsa.PrivateKey) uint64 {
	publicKey := sk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("error casting public key to ECDSA")
		return 0
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error("fetching nonce error ", err)
		return 0
	}
	return nonce
}

func (dmcTx *TransactionObj) SignTx() (*types.Transaction, error) {

	val := big.NewInt(dmcTx.Value)

	sk, err := crypto.HexToECDSA(dmcTx.PrivKey)
	if err != nil {
		log.Error("hex to ecdsa error ", err)
		return nil, err
	}
	nonce := calculatePendingNonce(sk)
	if nonce < 0 {
		return nil, errors.New("invalid nonce")
	}

	signer := types.NewEIP155Signer(chainId)

	dmcTx.GasPrice, err = ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("suggest gas price error: "+err.Error())
	}

	var baseTx *types.LegacyTx
	if len(dmcTx.To) != 0 {
		toAdd := common.HexToAddress(dmcTx.To)
		baseTx = &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: dmcTx.GasPrice,
			Gas:      dmcTx.Gas,
			To:       &toAdd,
			Value:    val,
			Data:     dmcTx.Data,
		}
	} else {
		baseTx = &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: dmcTx.GasPrice,
			Gas:      dmcTx.Gas,
			Value:    val,
			Data:     dmcTx.Data,
		}
	}
	tx := types.NewTx(baseTx)
	signedTx, _ := types.SignTx(tx, signer, sk)

	return signedTx, nil
}

func sendSignedTx(tx *types.Transaction) (string, error){

	err := ethClient.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Error("send transaction error ", err)
		return "", err
	}

	txHash := tx.Hash().String()
	log.Info("tx with hash ", txHash)
	return txHash, nil
}

func callContract(msg ethereum.CallMsg, blockNum *big.Int) (string, error) {
	res, err := ethClient.CallContract(context.Background(), msg, blockNum)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (dmcTx *TransactionObj) sendTx() (string, error) {
	SignedTx, err := dmcTx.SignTx()
	if err != nil || SignedTx == nil {
		return "", err
	}
	return sendSignedTx(SignedTx)
}

func (dmcTx *TransactionObj) deployContract() (string, error) {
	if dmcTx.To != "" {
		err := errors.New("contract creation does not need \"TO\" address")
		return "", err
	}
	return dmcTx.sendTx()
}

func getTransactionByHash(txHash string)  (*types.Transaction, error){

	hash := common.HexToHash(txHash)
	tx, _, err := ethClient.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Error("get transaction by hash error %v", err)
		return nil, err
	}
	return tx, nil
}


func getTransactionReceipt(txHash string) (*types.Receipt, error) {
	hash := common.HexToHash(txHash)
	txReceipt, err := ethClient.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Error("get Transaction Receipt error %v", err)
		return nil, err
	}
	return txReceipt, nil
}