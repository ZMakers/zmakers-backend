package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zmakers-backend/service"
)

func TestQueryUser(t *testing.T) {
	testUsers := [...]User{
		{
			Name: "ZJJ1",
			//Email: "zxczxcs@126.com",
			//Password: "ZZZXXX",
			//Hash: "ZZZXXX",
		},
	}
	service.DbInit()
	for _, user := range testUsers {
		//_ = user.GenerateToken()
		//err := user.CreatUser()
		//assert.Nil(t, err)
		user.SearchUser()
	}
}


func TestGenerateToken(t *testing.T)  {
	testUsers := [...]User{
		{
			Name: "ZJJ",
			Email: "zxczxc@126.com",
			Password: "ZZZXXX",
			Hash: "0x1",
		},
		{
			Name: "ZXX",
			Email: "AAA@126.com",
			Password: "123XXX",
			Hash: "0x2",
		},
		{
			Name: "ZXX",
			Email: "AAA@126.com",
			Password: "123XXX",
			Hash: "0x3",
		},
	}


	for _, user := range testUsers {
		err := user.CreateBefore()
		assert.Nil(t, err)
		err = user.GenerateToken()
		assert.Nil(t, err)
		user.VerifyToken(user.Token)
	}

}
