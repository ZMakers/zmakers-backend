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

func (m *Media) CreateMedia() int64 {
	result := service.Db.Table("medias").Create(m)
	if result.Error != nil && result.RowsAffected <= 0 {
		fmt.Println(result.Error)
		panic(result.Error)
	}
	return result.RowsAffected
}
