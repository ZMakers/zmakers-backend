package models

import "zmakers-backend/service"

type Tbl_Nft struct {
	TokenId uint64 `json:"tokenId" gorm:"column:tokenId"`
	TxHash string `json:"txHash" gorm:"column:txHash"`
	Address string `json:"address" gorm:"column:address"`
	TokenURI string `json:"tokenURI" gorm:"column:tokenURI"`
}

func (n *Tbl_Nft) CreatNft() error {
	result := service.Db.Create(n)
	if result.Error != nil || result.RowsAffected <= 0 {
		return result.Error
	}
	return nil
}

func (n *Tbl_Nft) GetNftByTxHash() *Tbl_Nft  {
	result := service.Db.Where("txHash = ?", n.TxHash).Find(n)
	if result.RowsAffected > 0 {
		return n
	}
	return nil
}