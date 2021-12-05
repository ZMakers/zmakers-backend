package models

import (
	"zmakers-backend/service"
)

type Collection struct {
	Owner    string `json:"owner"`
	MetaPath string `json:"metaPath" gorm:"column:metaPath"`
	TxHash string `json:"txHash" gorm:"column:txHash"`
}

func (c *Collection) PreCheck() bool {
	if c.Owner == "" || c.MetaPath == "" {
		return false
	}
	return true
}

func (c *Collection) CreateCollection() int64 {
	result := service.Db.Create(c)
	if result.Error != nil && result.RowsAffected <= 0 {
		panic(result.Error)
	}
	return result.RowsAffected
}

