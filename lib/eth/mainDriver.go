package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"zmakers-backend/lib/eth/nft"
)

func deployDigitalMediaCore(priKey string) {
	sk, _ := crypto.HexToECDSA(priKey)
	auth := bind.NewKeyedTransactor(sk)
	auth.Nonce = big.NewInt(int64(calculatePendingNonce(sk)))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(6123456) // in units
	auth.GasPrice = big.NewInt(2000000000)

	dmsAddress := common.HexToAddress("0xEB77CA33f7219edE12fbC8c8f282f06634F27B2e")
	crsAddress := common.HexToAddress("0x1D42B5f0a0a19E82D354706BE9Df863F9fcb2a33")
	address, tx, instance, err := nft.DeployDigitalMediaCore(auth, ethClient, "DMC", "ZJJ", big.NewInt(1), dmsAddress, crsAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance // will be using the instance in the next section

}

