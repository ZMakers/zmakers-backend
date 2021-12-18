package models

import (
	"fmt"
	"zmakers-backend/service"
)

type Media struct {
	Id uint64 `json:"id" gorm:"column:id"`
	TotalSupply uint32	`json:"totalSupply" gorm:"column:totalSupply"`
	CollectionId uint64 `json:"collectionId" gorm:"column:collectionId"`
	MetadataPath string	`json:"metadataPath" gorm:"column:metadataPath"`
	NumRelease uint32	`json:"numRelease" gorm:"column:numRelease"`
	TxHash string `json:"txHash" gorm:"column:txHash"`
}

type Transfer struct {
	TxHash string    `gorm:"column:txHash" json:"txHash" form:"txHash"`
	From string    `gorm:"column:from" json:"from" form:"from"`
	To string    `gorm:"column:to" json:"to" form:"to"`
	MediaId uint64    `gorm:"column:mediaId" json:"mediaId" form:"mediaId"`
}

func (m *Media) CreateMedia() int64 {
	result := service.Db.Table("medias").Create(m)
	if result.Error != nil && result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return result.RowsAffected
}

func (m *Media) CreateMediaReleases() int64 {
	result := service.Db.Table("medias").Create(m)
	if result.Error != nil && result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return result.RowsAffected
}

func (t *Transfer) BuyMedia() int64 {
	result := service.Db.Table("transfers").Create(t)
	if result.Error != nil && result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		panic(result.Error)
	}

	result = service.Db.Where("txHash = ?",t.TxHash).Find(t)
	if result.Error != nil && result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return result.RowsAffected
}

